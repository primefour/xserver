package model

import (
	"github.com/primefour/xserver/utils"
	"net/url"
)

type ServiceSettings struct {
	SiteURL                                  *string
	ListenAddress                            string
	ConnectionSecurity                       *string
	TLSCertFile                              *string
	TLSKeyFile                               *string
	UseLetsEncrypt                           *bool
	LetsEncryptCertificateCacheFile          *string
	Forward80To443                           *bool
	ReadTimeout                              *int
	WriteTimeout                             *int
	MaximumLoginAttempts                     int
	GoogleDeveloperKey                       string
	EnableOAuthServiceProvider               bool
	EnableIncomingWebhooks                   bool
	EnableOutgoingWebhooks                   bool
	EnableCommands                           *bool
	EnableOnlyAdminIntegrations              *bool
	EnablePostUsernameOverride               bool
	EnablePostIconOverride                   bool
	EnableLinkPreviews                       *bool
	EnableTesting                            bool
	EnableDeveloper                          *bool
	EnableSecurityFixAlert                   *bool
	EnableInsecureOutgoingConnections        *bool
	EnableMultifactorAuthentication          *bool
	EnforceMultifactorAuthentication         *bool
	AllowCorsFrom                            *string
	SessionLengthWebInDays                   *int
	SessionLengthMobileInDays                *int
	SessionLengthSSOInDays                   *int
	SessionCacheInMinutes                    *int
	WebsocketSecurePort                      *int
	WebsocketPort                            *int
	WebserverMode                            *string
	EnableCustomEmoji                        *bool
	RestrictCustomEmojiCreation              *string
	RestrictPostDelete                       *string
	AllowEditPost                            *string
	PostEditTimeLimit                        *int
	TimeBetweenUserTypingUpdatesMilliseconds *int64
	EnablePostSearch                         *bool
	EnableUserTypingMessages                 *bool
	EnableUserStatuses                       *bool
	ClusterLogTimeoutMilliseconds            *int
}

func (self *ServiceSettings) SetDefault() {

	if self.SiteURL == nil {
		self.SiteURL = new(string)
		*self.SiteURL = SERVICE_SETTINGS_DEFAULT_SITE_URL
	}

	if self.EnableLinkPreviews == nil {
		self.EnableLinkPreviews = new(bool)
		*self.EnableLinkPreviews = false
	}

	if self.EnableDeveloper == nil {
		self.EnableDeveloper = new(bool)
		*self.EnableDeveloper = false
	}

	if self.EnableSecurityFixAlert == nil {
		self.EnableSecurityFixAlert = new(bool)
		*self.EnableSecurityFixAlert = true
	}

	if self.EnableInsecureOutgoingConnections == nil {
		self.EnableInsecureOutgoingConnections = new(bool)
		*self.EnableInsecureOutgoingConnections = false
	}

	if self.EnableMultifactorAuthentication == nil {
		self.EnableMultifactorAuthentication = new(bool)
		*self.EnableMultifactorAuthentication = false
	}

	if self.EnforceMultifactorAuthentication == nil {
		self.EnforceMultifactorAuthentication = new(bool)
		*self.EnforceMultifactorAuthentication = false
	}

	if self.SessionLengthWebInDays == nil {
		self.SessionLengthWebInDays = new(int)
		*self.SessionLengthWebInDays = 30
	}

	if self.SessionLengthMobileInDays == nil {
		self.SessionLengthMobileInDays = new(int)
		*self.SessionLengthMobileInDays = 30
	}

	if self.SessionLengthSSOInDays == nil {
		self.SessionLengthSSOInDays = new(int)
		*self.SessionLengthSSOInDays = 30
	}

	if self.SessionCacheInMinutes == nil {
		self.SessionCacheInMinutes = new(int)
		*self.SessionCacheInMinutes = 10
	}

	if self.EnableCommands == nil {
		self.EnableCommands = new(bool)
		*self.EnableCommands = false
	}

	if self.EnableOnlyAdminIntegrations == nil {
		self.EnableOnlyAdminIntegrations = new(bool)
		*self.EnableOnlyAdminIntegrations = true
	}

	if self.WebsocketPort == nil {
		self.WebsocketPort = new(int)
		*self.WebsocketPort = 80
	}

	if self.WebsocketSecurePort == nil {
		self.WebsocketSecurePort = new(int)
		*self.WebsocketSecurePort = 443
	}

	if self.AllowCorsFrom == nil {
		self.AllowCorsFrom = new(string)
		*self.AllowCorsFrom = SERVICE_SETTINGS_DEFAULT_ALLOW_CORS_FROM
	}

	if self.WebserverMode == nil {
		self.WebserverMode = new(string)
		*self.WebserverMode = "gzip"
	} else if *self.WebserverMode == "regular" {
		*self.WebserverMode = "gzip"
	}

	if self.EnableCustomEmoji == nil {
		self.EnableCustomEmoji = new(bool)
		*self.EnableCustomEmoji = true
	}

	if self.RestrictCustomEmojiCreation == nil {
		self.RestrictCustomEmojiCreation = new(string)
		*self.RestrictCustomEmojiCreation = RESTRICT_EMOJI_CREATION_ALL
	}

	if self.RestrictPostDelete == nil {
		self.RestrictPostDelete = new(string)
		*self.RestrictPostDelete = PERMISSIONS_DELETE_POST_ALL
	}

	if self.AllowEditPost == nil {
		self.AllowEditPost = new(string)
		*self.AllowEditPost = ALLOW_EDIT_POST_ALWAYS
	}

	if self.PostEditTimeLimit == nil {
		self.PostEditTimeLimit = new(int)
		*self.PostEditTimeLimit = 300
	}

	if self.TimeBetweenUserTypingUpdatesMilliseconds == nil {
		self.TimeBetweenUserTypingUpdatesMilliseconds = new(int64)
		*self.TimeBetweenUserTypingUpdatesMilliseconds = 5000
	}

	if self.EnablePostSearch == nil {
		self.EnablePostSearch = new(bool)
		*self.EnablePostSearch = true
	}

	if self.EnableUserTypingMessages == nil {
		self.EnableUserTypingMessages = new(bool)
		*self.EnableUserTypingMessages = true
	}

	if self.EnableUserStatuses == nil {
		self.EnableUserStatuses = new(bool)
		*self.EnableUserStatuses = true
	}

	if self.ClusterLogTimeoutMilliseconds == nil {
		self.ClusterLogTimeoutMilliseconds = new(int)
		*self.ClusterLogTimeoutMilliseconds = 2000
	}

	if self.ConnectionSecurity == nil {
		self.ConnectionSecurity = new(string)
		*self.ConnectionSecurity = ""
	}

	if self.TLSKeyFile == nil {
		self.TLSKeyFile = new(string)
		*self.TLSKeyFile = SERVICE_SETTINGS_DEFAULT_TLS_KEY_FILE
	}

	if self.TLSCertFile == nil {
		self.TLSCertFile = new(string)
		*self.TLSCertFile = SERVICE_SETTINGS_DEFAULT_TLS_CERT_FILE
	}

	if self.UseLetsEncrypt == nil {
		self.UseLetsEncrypt = new(bool)
		*self.UseLetsEncrypt = false
	}

	if self.LetsEncryptCertificateCacheFile == nil {
		self.LetsEncryptCertificateCacheFile = new(string)
		*self.LetsEncryptCertificateCacheFile = "./config/letsencrypt.cache"
	}

	if self.ReadTimeout == nil {
		self.ReadTimeout = new(int)
		*self.ReadTimeout = SERVICE_SETTINGS_DEFAULT_READ_TIMEOUT
	}

	if self.WriteTimeout == nil {
		self.WriteTimeout = new(int)
		*self.WriteTimeout = SERVICE_SETTINGS_DEFAULT_WRITE_TIMEOUT
	}

	if self.Forward80To443 == nil {
		self.Forward80To443 = new(bool)
		*self.Forward80To443 = false
	}
}

func (self *ServiceSettings) IsValidate() *utils.AppError {

	if self.MaximumLoginAttempts <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.login_attempts.app_error", nil, "")
	}

	if len(*self.SiteURL) != 0 {
		if _, err := url.ParseRequestURI(*self.SiteURL); err != nil {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.site_url.app_error", nil, "")
		}
	}

	if len(self.ListenAddress) == 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.listen_address.app_error", nil, "")
	}

	if len(*self.SiteURL) == 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.site_url_email_batching.app_error", nil, "")
	}

	if !(*self.ConnectionSecurity == CONN_SECURITY_NONE || *self.ConnectionSecurity == CONN_SECURITY_TLS) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webserver_security.app_error", nil, "")
	}

	if *self.ReadTimeout <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.read_timeout.app_error", nil, "")
	}

	if *self.WriteTimeout <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.write_timeout.app_error", nil, "")
	}

	if *self.TimeBetweenUserTypingUpdatesMilliseconds < 1000 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.time_between_user_typing.app_error", nil, "")
	}
	return nil
}
