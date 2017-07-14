package api

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
	_ "github.com/nicksnyder/go-i18n/i18n"
	"github.com/primefour/xserver/app"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/utils"
	"net/http"
)

type Routes struct {
	Root    *mux.Router // ''
	ApiRoot *mux.Router // 'api/v1'

	Users          *mux.Router // 'api/v1/users'
	User           *mux.Router // 'api/v1/users/{user_id:[A-Za-z0-9]+}'
	UserByUsername *mux.Router // 'api/v1/users/username/{username:[A-Za-z0-9_-\.]+}'
	UserByEmail    *mux.Router // 'api/v1/users/email/{email}'

	OAuth     *mux.Router // 'api/v1/oauth'
	OAuthApps *mux.Router // 'api/v1/oauth/apps'
	OAuthApp  *mux.Router // 'api/v1/oauth/apps/{app_id:[A-Za-z0-9]+}'
	System    *mux.Router // 'api/v1/system'
}

const (
	CLIENT_DIR        = "/home/crazyhorse/CodeWork/GoWorkSpace/case/src/github.com/mattermost/platform/webapp/dist"
	API_URL_SUFFIX_V1 = "/api/v1"
	API_URL_SUFFIX    = API_URL_SUFFIX_V1
)

var BaseRoutes *Routes

func InitRouter() {
	app.Srv.Router = mux.NewRouter()
	app.Srv.Router.NotFoundHandler = http.HandlerFunc(Handle404)
}

func InitApi(full bool) {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = app.Srv.Router
	BaseRoutes.ApiRoot = app.Srv.Router.PathPrefix(API_URL_SUFFIX).Subrouter()

	BaseRoutes.Users = BaseRoutes.ApiRoot.PathPrefix("/users").Subrouter()
	BaseRoutes.User = BaseRoutes.ApiRoot.PathPrefix("/users/{user_id:[A-Za-z0-9]+}").Subrouter()
	BaseRoutes.UserByUsername = BaseRoutes.Users.PathPrefix("/username/{username:[A-Za-z0-9\\_\\-\\.]+}").Subrouter()
	BaseRoutes.UserByEmail = BaseRoutes.Users.PathPrefix("/email/{email}").Subrouter()

	BaseRoutes.OAuth = BaseRoutes.ApiRoot.PathPrefix("/oauth").Subrouter()
	BaseRoutes.OAuthApps = BaseRoutes.OAuth.PathPrefix("/apps").Subrouter()
	BaseRoutes.OAuthApp = BaseRoutes.OAuthApps.PathPrefix("/{app_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.System = BaseRoutes.ApiRoot.PathPrefix("/system").Subrouter()

	utils.InitHTML()
}

func HandleEtag(etag string, routeName string, w http.ResponseWriter, r *http.Request) bool {
	if et := r.Header.Get(utils.HEADER_ETAG_CLIENT); len(etag) > 0 {
		if et == etag {
			w.Header().Set(utils.HEADER_ETAG_SERVER, etag)
			w.WriteHeader(http.StatusNotModified)
			return true
		}
	}
	return false
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	err := model.NewLocAppError("Handle404", "api.context.404.app_error", nil, "")
	err.Translate(utils.T)
	err.StatusCode = http.StatusNotFound

	l4g.Debug("%v: code=404 ip=%v", r.URL.Path, utils.GetIpAddress(r))

	w.WriteHeader(err.StatusCode)
	err.DetailedError = "There doesn't appear to be an api call for the url='" + r.URL.Path + "'."
	w.Write([]byte(err.ToJson()))
}

func ReturnStatusOK(w http.ResponseWriter) {
}
