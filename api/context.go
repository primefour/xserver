package api

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	l4g "github.com/alecthomas/log4go"
	goi18n "github.com/nicksnyder/go-i18n/i18n"

	"github.com/primefour/xserver/app"
	"github.com/primefour/xserver/einterfaces"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/utils"
)

const (
	HEADER_REQUEST_ID         = "X-Request-ID"
	HEADER_VERSION_ID         = "X-Version-ID"
	HEADER_CLUSTER_ID         = "X-Cluster-ID"
	HEADER_ETAG_SERVER        = "ETag"
	HEADER_ETAG_CLIENT        = "If-None-Match"
	HEADER_FORWARDED          = "X-Forwarded-For"
	HEADER_REAL_IP            = "X-Real-IP"
	HEADER_FORWARDED_PROTO    = "X-Forwarded-Proto"
	HEADER_TOKEN              = "token"
	HEADER_BEARER             = "BEARER"
	HEADER_AUTH               = "Authorization"
	HEADER_REQUESTED_WITH     = "X-Requested-With"
	HEADER_REQUESTED_WITH_XML = "XMLHttpRequest"
)

type Context struct {
	Session       model.Session
	Params        *ApiParams
	Err           *model.AppError
	T             goi18n.TranslateFunc
	RequestId     string
	IpAddress     string
	Path          string
	siteURLHeader string
}

type handler struct {
	handleFunc     func(*Context, http.ResponseWriter, *http.Request)
	requireSession bool
	trustRequester bool
	requireMfa     bool
}

func AppHandlerIndependent(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, false, false}
}

func ApiHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:     h,
		requireSession: false,
		trustRequester: false,
		requireMfa:     false,
	}
}

func ApiSessionRequired(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:     h,
		requireSession: true,
		trustRequester: false,
		requireMfa:     true,
	}
}

func ApiSessionRequiredMfa(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:     h,
		requireSession: true,
		trustRequester: false,
		requireMfa:     false,
	}
}

func ApiHandlerTrustRequester(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:     h,
		requireSession: false,
		trustRequester: true,
		requireMfa:     false,
	}
}

func IsApiCall(r *http.Request) bool {
	return strings.Index(r.URL.Path, "/api/") == 0
}

func ApiSessionRequiredTrustRequester(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:     h,
		requireSession: true,
		trustRequester: true,
		requireMfa:     true,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	l4g.Debug("%v - %v", r.Method, r.URL.Path)

	c := &Context{}
	c.T, _ = utils.GetTranslationsAndLocale(w, r)
	c.RequestId = model.NewId()
	c.IpAddress = utils.GetIpAddress(r)
	c.Params = ApiParamsFromRequest(r)

	token := ""
	isTokenFromQueryString := false

	// Attempt to parse token out of the header
	authHeader := r.Header.Get(HEADER_AUTH)
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == HEADER_BEARER {
		// Default session token
		token = authHeader[7:]

	} else if len(authHeader) > 5 && strings.ToLower(authHeader[0:5]) == HEADER_TOKEN {
		// OAuth token
		token = authHeader[6:]
	}

	// Attempt to parse the token from the cookie
	if len(token) == 0 {
		if cookie, err := r.Cookie(model.SESSION_COOKIE_TOKEN); err == nil {
			token = cookie.Value

			if h.requireSession && !h.trustRequester {
				if r.Header.Get(HEADER_REQUESTED_WITH) != HEADER_REQUESTED_WITH_XML {
					c.Err = model.NewLocAppError("ServeHTTP", "api.context.session_expired.app_error", nil, "token="+token+" Appears to be a CSRF attempt")
					token = ""
				}
			}
		}
	}

	// Attempt to parse token out of the query string
	if len(token) == 0 {
		token = r.URL.Query().Get("access_token")
		if token != "" {
			isTokenFromQueryString = true
		}
	}

	c.SetSiteURLHeader(app.GetProtocol(r) + "://" + r.Host)

	w.Header().Set(HEADER_REQUEST_ID, c.RequestId)
	w.Header().Set(HEADER_VERSION_ID, fmt.Sprintf("%v.%v", model.CurrentVersion, utils.ClientCfgHash))
	if einterfaces.GetClusterInterface() != nil {
		w.Header().Set(model.HEADER_CLUSTER_ID, einterfaces.GetClusterInterface().GetClusterId())
	}

	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		w.Header().Set("Expires", "0")
	}

	if len(token) != 0 {
		session, err := app.GetSession(token)

		if err != nil {
			l4g.Error(utils.T("api.context.invalid_session.error"), err.Error())
			c.RemoveSessionCookie(w, r)
			if h.requireSession {
				c.Err = model.NewLocAppError("ServeHTTP", "api.context.session_expired.app_error", nil, "token="+token)
				c.Err.StatusCode = http.StatusUnauthorized
			}
		} else if !session.IsOAuth && isTokenFromQueryString {
			c.Err = model.NewLocAppError("ServeHTTP", "api.context.token_provided.app_error", nil, "token="+token)
			c.Err.StatusCode = http.StatusUnauthorized
		} else {
			c.Session = *session
		}
	}

	c.Path = r.URL.Path

	if c.Err == nil && h.requireSession {
		c.SessionRequired()
	}

	if c.Err == nil {
		h.handleFunc(c, w, r)
	}

	// Handle errors that have occured
	if c.Err != nil {
		c.Err.Translate(c.T)
		c.Err.RequestId = c.RequestId
		c.LogError(c.Err)
		c.Err.Where = r.URL.Path

		// Block out detailed error when not in developer mode
		if !*utils.Cfg.ServiceSettings.EnableDeveloper {
			c.Err.DetailedError = ""
		}

		w.WriteHeader(c.Err.StatusCode)
		w.Write([]byte(c.Err.ToJson()))

		if einterfaces.GetMetricsInterface() != nil {
			einterfaces.GetMetricsInterface().IncrementHttpError()
		}
	}

	if einterfaces.GetMetricsInterface() != nil {
		einterfaces.GetMetricsInterface().IncrementHttpRequest()

		if r.URL.Path != model.API_URL_SUFFIX+"/users/websocket" {
			elapsed := float64(time.Since(now)) / float64(time.Second)
			einterfaces.GetMetricsInterface().ObserveHttpRequestDuration(elapsed)
		}
	}
}

func (c *Context) LogAudit(extraInfo string) {
	audit := &model.Audit{UserId: c.Session.UserId, IpAddress: c.IpAddress, Action: c.Path, ExtraInfo: extraInfo, SessionId: c.Session.Id}
	if r := <-app.Srv.Store.Audit().Save(audit); r.Err != nil {
		c.LogError(r.Err)
	}
}

func (c *Context) LogAuditWithUserId(userId, extraInfo string) {

	if len(c.Session.UserId) > 0 {
		extraInfo = strings.TrimSpace(extraInfo + " session_user=" + c.Session.UserId)
	}

	audit := &model.Audit{UserId: userId, IpAddress: c.IpAddress, Action: c.Path, ExtraInfo: extraInfo, SessionId: c.Session.Id}
	if r := <-app.Srv.Store.Audit().Save(audit); r.Err != nil {
		c.LogError(r.Err)
	}
}

func (c *Context) LogError(err *model.AppError) {
	// filter out endless reconnects
	if c.Path == "/api/v1/users/websocket" && err.StatusCode == 401 || err.Id == "web.check_browser_compatibility.app_error" {
		c.LogDebug(err)
	} else {
		l4g.Error(utils.TDefault("api.context.log.error"), c.Path, err.Where, err.StatusCode,
			c.RequestId, c.Session.UserId, c.IpAddress, err.SystemMessage(utils.TDefault), err.DetailedError)
	}
}

func (c *Context) LogDebug(err *model.AppError) {
	l4g.Debug(utils.TDefault("api.context.log.error"), c.Path, err.Where, err.StatusCode,
		c.RequestId, c.Session.UserId, c.IpAddress, err.SystemMessage(utils.TDefault), err.DetailedError)
}

func (c *Context) IsSystemAdmin() bool {
	return app.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM)
}

func (c *Context) SessionRequired() {
	if len(c.Session.UserId) == 0 {
		c.Err = model.NewAppError("", "api.context.session_expired.app_error", nil, "UserRequired", http.StatusUnauthorized)
		return
	}
}

func (c *Context) RemoveSessionCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     model.SESSION_COOKIE_TOKEN,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}

func (c *Context) SetInvalidParam(parameter string) {
	c.Err = NewInvalidParamError(parameter)
}

func (c *Context) SetInvalidUrlParam(parameter string) {
	c.Err = NewInvalidUrlParamError(parameter)
}

func NewInvalidParamError(parameter string) *model.AppError {
	err := model.NewLocAppError("Context", "api.context.invalid_body_param.app_error", map[string]interface{}{"Name": parameter}, "")
	err.StatusCode = http.StatusBadRequest
	return err
}
func NewInvalidUrlParamError(parameter string) *model.AppError {
	err := model.NewLocAppError("Context", "api.context.invalid_url_param.app_error", map[string]interface{}{"Name": parameter}, "")
	err.StatusCode = http.StatusBadRequest
	return err
}

func (c *Context) SetPermissionError(permission *model.Permission) {
	c.Err = model.NewLocAppError("Permissions", "api.context.permissions.app_error", nil, "userId="+c.Session.UserId+", "+"permission="+permission.Id)
	c.Err.StatusCode = http.StatusForbidden
}

func (c *Context) SetSiteURLHeader(url string) {
	c.siteURLHeader = strings.TrimRight(url, "/")
}

func (c *Context) GetSiteURLHeader() string {
	return c.siteURLHeader
}

func (c *Context) RequireUserId() *Context {
	if c.Err != nil {
		return c
	}

	if c.Params.UserId == model.ME {
		c.Params.UserId = c.Session.UserId
	}

	if len(c.Params.UserId) != 26 {
		c.SetInvalidUrlParam("user_id")
	}
	return c
}

func (c *Context) RequireTeamId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.TeamId) != 26 {
		c.SetInvalidUrlParam("team_id")
	}
	return c
}

func (c *Context) RequireChannelId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.ChannelId) != 26 {
		c.SetInvalidUrlParam("channel_id")
	}
	return c
}

func (c *Context) RequireUsername() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsValidUsername(c.Params.Username) {
		c.SetInvalidParam("username")
	}

	return c
}

func (c *Context) RequirePostId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.PostId) != 26 {
		c.SetInvalidUrlParam("post_id")
	}
	return c
}

func (c *Context) RequireAppId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.AppId) != 26 {
		c.SetInvalidUrlParam("app_id")
	}
	return c
}

func (c *Context) RequireFileId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.FileId) != 26 {
		c.SetInvalidUrlParam("file_id")
	}

	return c
}

func (c *Context) RequireReportId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.ReportId) != 26 {
		c.SetInvalidUrlParam("report_id")
	}
	return c
}

func (c *Context) RequireEmojiId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.EmojiId) != 26 {
		c.SetInvalidUrlParam("emoji_id")
	}
	return c
}

func (c *Context) RequireTeamName() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsValidTeamName(c.Params.TeamName) {
		c.SetInvalidUrlParam("team_name")
	}

	return c
}

func (c *Context) RequireChannelName() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsValidChannelIdentifier(c.Params.ChannelName) {
		c.SetInvalidUrlParam("channel_name")
	}

	return c
}

func (c *Context) RequireEmail() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsValidEmail(c.Params.Email) {
		c.SetInvalidUrlParam("email")
	}

	return c
}

func (c *Context) RequireCategory() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsValidAlphaNumHyphenUnderscore(c.Params.Category, true) {
		c.SetInvalidUrlParam("category")
	}

	return c
}

func (c *Context) RequireService() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.Service) == 0 {
		c.SetInvalidUrlParam("service")
	}

	return c
}

func (c *Context) RequirePreferenceName() *Context {
	if c.Err != nil {
		return c
	}

	if !model.IsValidAlphaNumHyphenUnderscore(c.Params.PreferenceName, true) {
		c.SetInvalidUrlParam("preference_name")
	}

	return c
}

func (c *Context) RequireEmojiName() *Context {
	if c.Err != nil {
		return c
	}

	validName := regexp.MustCompile(`^[a-zA-Z0-9\-\+_]+$`)

	if len(c.Params.EmojiName) == 0 || len(c.Params.EmojiName) > 64 || !validName.MatchString(c.Params.EmojiName) {
		c.SetInvalidUrlParam("emoji_name")
	}

	return c
}

func (c *Context) RequireHookId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.HookId) != 26 {
		c.SetInvalidUrlParam("hook_id")
	}

	return c
}

func (c *Context) RequireCommandId() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.CommandId) != 26 {
		c.SetInvalidUrlParam("command_id")
	}
	return c
}
