package model

type ServiceSettings struct {
	SiteURL                                           *string
	LicenseFileLocation                               *string
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
	EnableAPIv3                                       *bool
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

	if s.LicenseFileLocation == nil {
		s.LicenseFileLocation = NewString("")
	}

	if s.ListenAddress == nil {
		s.ListenAddress = NewString(SERVICE_SETTINGS_DEFAULT_LISTEN_AND_ADDRESS)
	}

	if s.EnableAPIv3 == nil {
		s.EnableAPIv3 = NewBool(true)
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

type ClusterSettings struct {
	Enable                *bool
	ClusterName           *string
	OverrideHostname      *string
	UseIpAddress          *bool
	UseExperimentalGossip *bool
	ReadOnlyConfig        *bool
	GossipPort            *int
	StreamingPort         *int
}

func (s *ClusterSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.ClusterName == nil {
		s.ClusterName = NewString("")
	}

	if s.OverrideHostname == nil {
		s.OverrideHostname = NewString("")
	}

	if s.UseIpAddress == nil {
		s.UseIpAddress = NewBool(true)
	}

	if s.UseExperimentalGossip == nil {
		s.UseExperimentalGossip = NewBool(false)
	}

	if s.ReadOnlyConfig == nil {
		s.ReadOnlyConfig = NewBool(true)
	}

	if s.GossipPort == nil {
		s.GossipPort = NewInt(8074)
	}

	if s.StreamingPort == nil {
		s.StreamingPort = NewInt(8075)
	}
}

type MetricsSettings struct {
	Enable           *bool
	BlockProfileRate *int
	ListenAddress    *string
}

func (s *MetricsSettings) SetDefaults() {
	if s.ListenAddress == nil {
		s.ListenAddress = NewString(":8067")
	}

	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.BlockProfileRate == nil {
		s.BlockProfileRate = NewInt(0)
	}
}

type LogSettings struct {
	EnableConsole          bool
	ConsoleLevel           string
	EnableFile             bool
	FileLevel              string
	FileFormat             string
	FileLocation           string
	EnableWebhookDebugging bool
	EnableDiagnostics      *bool
}

func (s *LogSettings) SetDefaults() {
	if s.EnableDiagnostics == nil {
		s.EnableDiagnostics = NewBool(true)
	}
}

type FileSettings struct {
	EnableFileAttachments   *bool
	EnableMobileUpload      *bool
	EnableMobileDownload    *bool
	MaxFileSize             *int64
	DriverName              *string
	Directory               string
	EnablePublicLink        bool
	PublicLinkSalt          *string
	InitialFont             string
	AmazonS3AccessKeyId     string
	AmazonS3SecretAccessKey string
	AmazonS3Bucket          string
	AmazonS3Region          string
	AmazonS3Endpoint        string
	AmazonS3SSL             *bool
	AmazonS3SignV2          *bool
	AmazonS3SSE             *bool
	AmazonS3Trace           *bool
}

func (s *FileSettings) SetDefaults() {
	if s.DriverName == nil {
		s.DriverName = NewString(IMAGE_DRIVER_LOCAL)
	}

	if s.AmazonS3Endpoint == "" {
		// Defaults to "s3.amazonaws.com"
		s.AmazonS3Endpoint = "s3.amazonaws.com"
	}

	if s.AmazonS3SSL == nil {
		s.AmazonS3SSL = NewBool(true) // Secure by default.
	}

	if s.AmazonS3SignV2 == nil {
		s.AmazonS3SignV2 = new(bool)
		// Signature v2 is not enabled by default.
	}

	if s.AmazonS3SSE == nil {
		s.AmazonS3SSE = NewBool(false) // Not Encrypted by default.
	}

	if s.AmazonS3Trace == nil {
		s.AmazonS3Trace = NewBool(false)
	}

	if s.EnableFileAttachments == nil {
		s.EnableFileAttachments = NewBool(true)
	}

	if s.EnableMobileUpload == nil {
		s.EnableMobileUpload = NewBool(true)
	}

	if s.EnableMobileDownload == nil {
		s.EnableMobileDownload = NewBool(true)
	}

	if s.MaxFileSize == nil {
		s.MaxFileSize = NewInt64(52428800) // 50 MB
	}

	if s.PublicLinkSalt == nil || len(*s.PublicLinkSalt) == 0 {
		s.PublicLinkSalt = NewString(NewRandomString(32))
	}

	if s.InitialFont == "" {
		// Defaults to "luximbi.ttf"
		s.InitialFont = "luximbi.ttf"
	}

	if s.Directory == "" {
		s.Directory = "./data/"
	}
}
