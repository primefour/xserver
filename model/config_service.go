package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	SERVICE_SETTINGS_DEFAULT_SITE_URL           = ""
	SERVICE_SETTINGS_DEFAULT_TLS_CERT_FILE      = ""
	SERVICE_SETTINGS_DEFAULT_TLS_KEY_FILE       = ""
	SERVICE_SETTINGS_DEFAULT_READ_TIMEOUT       = 300
	SERVICE_SETTINGS_DEFAULT_WRITE_TIMEOUT      = 300
	SERVICE_SETTINGS_DEFAULT_MAX_LOGIN_ATTEMPTS = 10
	SERVICE_SETTINGS_DEFAULT_ALLOW_CORS_FROM    = ""
	SERVICE_SETTINGS_DEFAULT_LISTEN_AND_ADDRESS = ":8065"

	//emoji permission
	RESTRICT_EMOJI_CREATION_ALL          = "all"
	RESTRICT_EMOJI_CREATION_ADMIN        = "admin"
	RESTRICT_EMOJI_CREATION_SYSTEM_ADMIN = "system_admin"

	//delete post permission
	PERMISSIONS_DELETE_POST_ALL          = "all"
	PERMISSIONS_DELETE_POST_TEAM_ADMIN   = "team_admin"
	PERMISSIONS_DELETE_POST_SYSTEM_ADMIN = "system_admin"

	//modify post permission
	ALLOW_EDIT_POST_ALWAYS     = "always"
	ALLOW_EDIT_POST_NEVER      = "never"
	ALLOW_EDIT_POST_TIME_LIMIT = "time_limit"

	//service config
	SERVICE_CONFIG_FILE_PATH = "./config/service_config.json"
	SERVICE_CONFIG_NAME      = "SERVICE_SETTINGS"
)

type ServiceSettings struct {
	SiteURL                                           *string
	ListenAddress                                     *string
	ConnectionSecurity                                *string
	TLSCertFile                                       *string
	TLSKeyFile                                        *string
	UseLetsEncrypt                                    *bool
	LetsEncryptCertificateCacheFile                   *string
	Forward80To443                                    *bool
	ReadTimeout                                       *int
	WriteTimeout                                      *int
	MaximumLoginAttempts                              *int
	GoroutineHealthThreshold                          *int
	GoogleDeveloperKey                                string
	EnableOAuthServiceProvider                        bool
	EnableIncomingWebhooks                            bool
	EnableOutgoingWebhooks                            bool
	EnableCommands                                    *bool
	EnableOnlyAdminIntegrations                       *bool
	EnablePostUsernameOverride                        bool
	EnablePostIconOverride                            bool
	EnableLinkPreviews                                *bool
	EnableTesting                                     bool
	EnableDeveloper                                   *bool
	EnableSecurityFixAlert                            *bool
	EnableInsecureOutgoingConnections                 *bool
	AllowedUntrustedInternalConnections               *string
	EnableMultifactorAuthentication                   *bool
	EnforceMultifactorAuthentication                  *bool
	EnableUserAccessTokens                            *bool
	AllowCorsFrom                                     *string
	SessionLengthWebInDays                            *int
	SessionLengthMobileInDays                         *int
	SessionLengthSSOInDays                            *int
	SessionCacheInMinutes                             *int
	SessionIdleTimeoutInMinutes                       *int
	WebsocketSecurePort                               *int
	WebsocketPort                                     *int
	WebserverMode                                     *string
	EnableCustomEmoji                                 *bool
	EnableEmojiPicker                                 *bool
	RestrictCustomEmojiCreation                       *string
	RestrictPostDelete                                *string
	AllowEditPost                                     *string
	PostEditTimeLimit                                 *int
	TimeBetweenUserTypingUpdatesMilliseconds          *int64
	EnablePostSearch                                  *bool
	EnableUserTypingMessages                          *bool
	EnableChannelViewedMessages                       *bool
	EnableUserStatuses                                *bool
	ExperimentalEnableAuthenticationTransfer          *bool
	ClusterLogTimeoutMilliseconds                     *int
	CloseUnusedDirectMessages                         *bool
	EnablePreviewFeatures                             *bool
	EnableTutorial                                    *bool
	ExperimentalEnableDefaultChannelLeaveJoinMessages *bool
	ExperimentalGroupUnreadChannels                   *bool
	ImageProxyType                                    *string
	ImageProxyURL                                     *string
	ImageProxyOptions                                 *string
}

func (s *ServiceSettings) SetDefaults() {
	if s.SiteURL == nil {
		s.SiteURL = NewString(SERVICE_SETTINGS_DEFAULT_SITE_URL)
	}

	if s.ListenAddress == nil {
		s.ListenAddress = NewString(SERVICE_SETTINGS_DEFAULT_LISTEN_AND_ADDRESS)
	}

	if s.EnableLinkPreviews == nil {
		s.EnableLinkPreviews = NewBool(false)
	}

	if s.EnableDeveloper == nil {
		s.EnableDeveloper = NewBool(false)
	}

	if s.EnableSecurityFixAlert == nil {
		s.EnableSecurityFixAlert = NewBool(true)
	}

	if s.EnableInsecureOutgoingConnections == nil {
		s.EnableInsecureOutgoingConnections = NewBool(false)
	}

	if s.AllowedUntrustedInternalConnections == nil {
		s.AllowedUntrustedInternalConnections = NewString("")
	}

	if s.EnableMultifactorAuthentication == nil {
		s.EnableMultifactorAuthentication = NewBool(false)
	}

	if s.EnforceMultifactorAuthentication == nil {
		s.EnforceMultifactorAuthentication = NewBool(false)
	}

	if s.EnableUserAccessTokens == nil {
		s.EnableUserAccessTokens = NewBool(false)
	}

	if s.GoroutineHealthThreshold == nil {
		s.GoroutineHealthThreshold = NewInt(-1)
	}

	if s.ConnectionSecurity == nil {
		s.ConnectionSecurity = NewString("")
	}

	if s.TLSKeyFile == nil {
		s.TLSKeyFile = NewString(SERVICE_SETTINGS_DEFAULT_TLS_KEY_FILE)
	}

	if s.TLSCertFile == nil {
		s.TLSCertFile = NewString(SERVICE_SETTINGS_DEFAULT_TLS_CERT_FILE)
	}

	if s.UseLetsEncrypt == nil {
		s.UseLetsEncrypt = NewBool(false)
	}

	if s.LetsEncryptCertificateCacheFile == nil {
		s.LetsEncryptCertificateCacheFile = NewString("./config/letsencrypt.cache")
	}

	if s.ReadTimeout == nil {
		s.ReadTimeout = NewInt(SERVICE_SETTINGS_DEFAULT_READ_TIMEOUT)
	}

	if s.WriteTimeout == nil {
		s.WriteTimeout = NewInt(SERVICE_SETTINGS_DEFAULT_WRITE_TIMEOUT)
	}

	if s.MaximumLoginAttempts == nil {
		s.MaximumLoginAttempts = NewInt(SERVICE_SETTINGS_DEFAULT_MAX_LOGIN_ATTEMPTS)
	}

	if s.Forward80To443 == nil {
		s.Forward80To443 = NewBool(false)
	}

	if s.TimeBetweenUserTypingUpdatesMilliseconds == nil {
		s.TimeBetweenUserTypingUpdatesMilliseconds = NewInt64(5000)
	}

	if s.EnablePostSearch == nil {
		s.EnablePostSearch = NewBool(true)
	}

	if s.EnableUserTypingMessages == nil {
		s.EnableUserTypingMessages = NewBool(true)
	}

	if s.EnableChannelViewedMessages == nil {
		s.EnableChannelViewedMessages = NewBool(true)
	}

	if s.EnableUserStatuses == nil {
		s.EnableUserStatuses = NewBool(true)
	}

	if s.ClusterLogTimeoutMilliseconds == nil {
		s.ClusterLogTimeoutMilliseconds = NewInt(2000)
	}

	if s.CloseUnusedDirectMessages == nil {
		s.CloseUnusedDirectMessages = NewBool(false)
	}

	if s.EnableTutorial == nil {
		s.EnableTutorial = NewBool(true)
	}

	if s.SessionLengthWebInDays == nil {
		s.SessionLengthWebInDays = NewInt(30)
	}

	if s.SessionLengthMobileInDays == nil {
		s.SessionLengthMobileInDays = NewInt(30)
	}

	if s.SessionLengthSSOInDays == nil {
		s.SessionLengthSSOInDays = NewInt(30)
	}

	if s.SessionCacheInMinutes == nil {
		s.SessionCacheInMinutes = NewInt(10)
	}

	if s.SessionIdleTimeoutInMinutes == nil {
		s.SessionIdleTimeoutInMinutes = NewInt(0)
	}

	if s.EnableCommands == nil {
		s.EnableCommands = NewBool(false)
	}

	if s.EnableOnlyAdminIntegrations == nil {
		s.EnableOnlyAdminIntegrations = NewBool(true)
	}

	if s.WebsocketPort == nil {
		s.WebsocketPort = NewInt(80)
	}

	if s.WebsocketSecurePort == nil {
		s.WebsocketSecurePort = NewInt(443)
	}

	//Access-Control-Allow-Origin
	if s.AllowCorsFrom == nil {
		s.AllowCorsFrom = NewString(SERVICE_SETTINGS_DEFAULT_ALLOW_CORS_FROM)
	}

	if s.WebserverMode == nil {
		s.WebserverMode = NewString("gzip")
	} else if *s.WebserverMode == "regular" {
		*s.WebserverMode = "gzip"
	}

	if s.EnableCustomEmoji == nil {
		s.EnableCustomEmoji = NewBool(false)
	}

	if s.EnableEmojiPicker == nil {
		s.EnableEmojiPicker = NewBool(true)
	}

	if s.RestrictCustomEmojiCreation == nil {
		s.RestrictCustomEmojiCreation = NewString(RESTRICT_EMOJI_CREATION_ALL)
	}

	if s.RestrictPostDelete == nil {
		s.RestrictPostDelete = NewString(PERMISSIONS_DELETE_POST_ALL)
	}

	if s.AllowEditPost == nil {
		s.AllowEditPost = NewString(ALLOW_EDIT_POST_ALWAYS)
	}

	if s.ExperimentalEnableAuthenticationTransfer == nil {
		s.ExperimentalEnableAuthenticationTransfer = NewBool(true)
	}

	if s.PostEditTimeLimit == nil {
		s.PostEditTimeLimit = NewInt(300)
	}

	if s.EnablePreviewFeatures == nil {
		s.EnablePreviewFeatures = NewBool(true)
	}

	if s.ExperimentalEnableDefaultChannelLeaveJoinMessages == nil {
		s.ExperimentalEnableDefaultChannelLeaveJoinMessages = NewBool(true)
	}

	if s.ExperimentalGroupUnreadChannels == nil {
		s.ExperimentalGroupUnreadChannels = NewBool(false)
	}

	if s.ImageProxyType == nil {
		s.ImageProxyType = NewString("")
	}

	if s.ImageProxyURL == nil {
		s.ImageProxyURL = NewString("")
	}

	if s.ImageProxyOptions == nil {
		s.ImageProxyOptions = NewString("")
	}
}

func servConfigParser(f *os.File) (interface{}, error) {
	settings := &ServiceSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("system settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetServSettings() *ServiceSettings {
	settings := utils.GetSettings(SERVICE_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*ServiceSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(SERVICE_CONFIG_NAME, SERVICE_CONFIG_FILE_PATH, true, servConfigParser)
	if err != nil {
		return
	}
}
