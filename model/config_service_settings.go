package model

import (
	"encoding/json"
	"github.com/primefour/xserver/utils"
	"io"
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

func (self *ServiceSettings) setDefault() {

	if self.ServiceSettings.SiteURL == nil {
		self.ServiceSettings.SiteURL = new(string)
		*self.ServiceSettings.SiteURL = SERVICE_SETTINGS_DEFAULT_SITE_URL
	}

	if self.ServiceSettings.EnableLinkPreviews == nil {
		self.ServiceSettings.EnableLinkPreviews = new(bool)
		*self.ServiceSettings.EnableLinkPreviews = false
	}

	if self.ServiceSettings.EnableDeveloper == nil {
		self.ServiceSettings.EnableDeveloper = new(bool)
		*self.ServiceSettings.EnableDeveloper = false
	}

	if self.ServiceSettings.EnableSecurityFixAlert == nil {
		self.ServiceSettings.EnableSecurityFixAlert = new(bool)
		*self.ServiceSettings.EnableSecurityFixAlert = true
	}

	if self.ServiceSettings.EnableInsecureOutgoingConnections == nil {
		self.ServiceSettings.EnableInsecureOutgoingConnections = new(bool)
		*self.ServiceSettings.EnableInsecureOutgoingConnections = false
	}

	if self.ServiceSettings.EnableMultifactorAuthentication == nil {
		self.ServiceSettings.EnableMultifactorAuthentication = new(bool)
		*self.ServiceSettings.EnableMultifactorAuthentication = false
	}

	if self.ServiceSettings.EnforceMultifactorAuthentication == nil {
		self.ServiceSettings.EnforceMultifactorAuthentication = new(bool)
		*self.ServiceSettings.EnforceMultifactorAuthentication = false
	}

	if self.ServiceSettings.SessionLengthWebInDays == nil {
		self.ServiceSettings.SessionLengthWebInDays = new(int)
		*self.ServiceSettings.SessionLengthWebInDays = 30
	}

	if self.ServiceSettings.SessionLengthMobileInDays == nil {
		self.ServiceSettings.SessionLengthMobileInDays = new(int)
		*self.ServiceSettings.SessionLengthMobileInDays = 30
	}

	if self.ServiceSettings.SessionLengthSSOInDays == nil {
		self.ServiceSettings.SessionLengthSSOInDays = new(int)
		*self.ServiceSettings.SessionLengthSSOInDays = 30
	}

	if self.ServiceSettings.SessionCacheInMinutes == nil {
		self.ServiceSettings.SessionCacheInMinutes = new(int)
		*self.ServiceSettings.SessionCacheInMinutes = 10
	}

	if self.ServiceSettings.EnableCommands == nil {
		self.ServiceSettings.EnableCommands = new(bool)
		*self.ServiceSettings.EnableCommands = false
	}

	if self.ServiceSettings.EnableOnlyAdminIntegrations == nil {
		self.ServiceSettings.EnableOnlyAdminIntegrations = new(bool)
		*self.ServiceSettings.EnableOnlyAdminIntegrations = true
	}

	if self.ServiceSettings.WebsocketPort == nil {
		self.ServiceSettings.WebsocketPort = new(int)
		*self.ServiceSettings.WebsocketPort = 80
	}

	if self.ServiceSettings.WebsocketSecurePort == nil {
		self.ServiceSettings.WebsocketSecurePort = new(int)
		*self.ServiceSettings.WebsocketSecurePort = 443
	}

	if self.ServiceSettings.AllowCorsFrom == nil {
		self.ServiceSettings.AllowCorsFrom = new(string)
		*self.ServiceSettings.AllowCorsFrom = SERVICE_SETTINGS_DEFAULT_ALLOW_CORS_FROM
	}

	if self.ServiceSettings.WebserverMode == nil {
		self.ServiceSettings.WebserverMode = new(string)
		*self.ServiceSettings.WebserverMode = "gzip"
	} else if *self.ServiceSettings.WebserverMode == "regular" {
		*self.ServiceSettings.WebserverMode = "gzip"
	}

	if self.ServiceSettings.EnableCustomEmoji == nil {
		self.ServiceSettings.EnableCustomEmoji = new(bool)
		*self.ServiceSettings.EnableCustomEmoji = true
	}

	if self.ServiceSettings.RestrictCustomEmojiCreation == nil {
		self.ServiceSettings.RestrictCustomEmojiCreation = new(string)
		*self.ServiceSettings.RestrictCustomEmojiCreation = RESTRICT_EMOJI_CREATION_ALL
	}

	if self.ServiceSettings.RestrictPostDelete == nil {
		self.ServiceSettings.RestrictPostDelete = new(string)
		*self.ServiceSettings.RestrictPostDelete = PERMISSIONS_DELETE_POST_ALL
	}

	if self.ServiceSettings.AllowEditPost == nil {
		self.ServiceSettings.AllowEditPost = new(string)
		*self.ServiceSettings.AllowEditPost = ALLOW_EDIT_POST_ALWAYS
	}

	if self.ServiceSettings.PostEditTimeLimit == nil {
		self.ServiceSettings.PostEditTimeLimit = new(int)
		*self.ServiceSettings.PostEditTimeLimit = 300
	}

	if self.ServiceSettings.TimeBetweenUserTypingUpdatesMilliseconds == nil {
		self.ServiceSettings.TimeBetweenUserTypingUpdatesMilliseconds = new(int64)
		*self.ServiceSettings.TimeBetweenUserTypingUpdatesMilliseconds = 5000
	}

	if self.ServiceSettings.EnablePostSearch == nil {
		self.ServiceSettings.EnablePostSearch = new(bool)
		*self.ServiceSettings.EnablePostSearch = true
	}

	if self.ServiceSettings.EnableUserTypingMessages == nil {
		self.ServiceSettings.EnableUserTypingMessages = new(bool)
		*self.ServiceSettings.EnableUserTypingMessages = true
	}

	if self.ServiceSettings.EnableUserStatuses == nil {
		self.ServiceSettings.EnableUserStatuses = new(bool)
		*self.ServiceSettings.EnableUserStatuses = true
	}

	if self.ServiceSettings.ClusterLogTimeoutMilliseconds == nil {
		self.ServiceSettings.ClusterLogTimeoutMilliseconds = new(int)
		*self.ServiceSettings.ClusterLogTimeoutMilliseconds = 2000
	}

	if self.ServiceSettings.ConnectionSecurity == nil {
		self.ServiceSettings.ConnectionSecurity = new(string)
		*self.ServiceSettings.ConnectionSecurity = ""
	}

	if self.ServiceSettings.TLSKeyFile == nil {
		self.ServiceSettings.TLSKeyFile = new(string)
		*self.ServiceSettings.TLSKeyFile = SERVICE_SETTINGS_DEFAULT_TLS_KEY_FILE
	}

	if self.ServiceSettings.TLSCertFile == nil {
		self.ServiceSettings.TLSCertFile = new(string)
		*self.ServiceSettings.TLSCertFile = SERVICE_SETTINGS_DEFAULT_TLS_CERT_FILE
	}

	if self.ServiceSettings.UseLetsEncrypt == nil {
		self.ServiceSettings.UseLetsEncrypt = new(bool)
		*self.ServiceSettings.UseLetsEncrypt = false
	}

	if self.ServiceSettings.LetsEncryptCertificateCacheFile == nil {
		self.ServiceSettings.LetsEncryptCertificateCacheFile = new(string)
		*self.ServiceSettings.LetsEncryptCertificateCacheFile = "./config/letsencrypt.cache"
	}

	if self.ServiceSettings.ReadTimeout == nil {
		self.ServiceSettings.ReadTimeout = new(int)
		*self.ServiceSettings.ReadTimeout = SERVICE_SETTINGS_DEFAULT_READ_TIMEOUT
	}

	if self.ServiceSettings.WriteTimeout == nil {
		self.ServiceSettings.WriteTimeout = new(int)
		*self.ServiceSettings.WriteTimeout = SERVICE_SETTINGS_DEFAULT_WRITE_TIMEOUT
	}

	if self.ServiceSettings.Forward80To443 == nil {
		self.ServiceSettings.Forward80To443 = new(bool)
		*self.ServiceSettings.Forward80To443 = false
	}
}

func (self *ServiceSettings) IsValidate() utils.AppError {

	if o.ServiceSettings.MaximumLoginAttempts <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.login_attempts.app_error", nil, "")
	}

	if len(*o.ServiceSettings.SiteURL) != 0 {
		if _, err := url.ParseRequestURI(*o.ServiceSettings.SiteURL); err != nil {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.site_url.app_error", nil, "")
		}
	}

	if len(o.ServiceSettings.ListenAddress) == 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.listen_address.app_error", nil, "")
	}

	if len(*o.ServiceSettings.SiteURL) == 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.site_url_email_batching.app_error", nil, "")
	}

	if !(*o.ServiceSettings.ConnectionSecurity == CONN_SECURITY_NONE || *o.ServiceSettings.ConnectionSecurity == CONN_SECURITY_TLS) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webserver_security.app_error", nil, "")
	}

	if *o.ServiceSettings.ReadTimeout <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.read_timeout.app_error", nil, "")
	}

	if *o.ServiceSettings.WriteTimeout <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.write_timeout.app_error", nil, "")
	}

	if *o.ServiceSettings.TimeBetweenUserTypingUpdatesMilliseconds < 1000 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.time_between_user_typing.app_error", nil, "")
	}
}
