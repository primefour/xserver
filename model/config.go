package model

import (
	"encoding/json"
	"github.com/primefour/xserver/utils"
	"io"
	"net/url"
)

type ClusterSettings struct {
	Enable                 *bool
	InterNodeListenAddress *string
	InterNodeUrls          []string
}

type MetricsSettings struct {
	Enable           *bool
	BlockProfileRate *int
	ListenAddress    *string
}

type AnalyticsSettings struct {
	MaxUsersForStatistics *int
}

type ComplianceSettings struct {
	Enable      *bool
	Directory   *string
	EnableDaily *bool
}

type LocalizationSettings struct {
	DefaultServerLocale *string
	DefaultClientLocale *string
	AvailableLocales    *string
}

type NativeAppSettings struct {
	AppDownloadLink        *string
	AndroidAppDownloadLink *string
	IosAppDownloadLink     *string
}

type Config struct {
	ServiceSettings      ServiceSettings
	TeamSettings         TeamSettings
	SqlSettings          SqlSettings
	LogSettings          LogSettings
	PasswordSettings     PasswordSettings
	FileSettings         FileSettings
	EmailSettings        EmailSettings
	RateLimitSettings    RateLimitSettings
	PrivacySettings      PrivacySettings
	SupportSettings      SupportSettings
	GitLabSettings       SSOSettings
	GoogleSettings       SSOSettings
	Office365Settings    SSOSettings
	LdapSettings         LdapSettings
	ComplianceSettings   ComplianceSettings
	LocalizationSettings LocalizationSettings
	SamlSettings         SamlSettings
	NativeAppSettings    NativeAppSettings
	ClusterSettings      ClusterSettings
	MetricsSettings      MetricsSettings
	AnalyticsSettings    AnalyticsSettings
	WebrtcSettings       WebrtcSettings
}

func (o *Config) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o *Config) GetSSOService(service string) *SSOSettings {
	switch service {
	case SERVICE_GITLAB:
		return &o.GitLabSettings
	case SERVICE_GOOGLE:
		return &o.GoogleSettings
	case SERVICE_OFFICE365:
		return &o.Office365Settings
	}

	return nil
}

func ConfigFromJson(data io.Reader) *Config {
	decoder := json.NewDecoder(data)
	var o Config
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *Config) SetDefaults() {

	if len(o.SqlSettings.AtRestEncryptKey) == 0 {
		o.SqlSettings.AtRestEncryptKey = NewRandomString(32)
	}

	if o.FileSettings.AmazonS3Endpoint == "" {
		// Defaults to "s3.amazonaws.com"
		o.FileSettings.AmazonS3Endpoint = "s3.amazonaws.com"
	}

	if o.FileSettings.AmazonS3Region == "" {
		// Defaults to "us-east-1" region.
		o.FileSettings.AmazonS3Region = "us-east-1"
	}

	if o.FileSettings.AmazonS3SSL == nil {
		o.FileSettings.AmazonS3SSL = new(bool)
		*o.FileSettings.AmazonS3SSL = true // Secure by default.
	}

	if o.FileSettings.EnableFileAttachments == nil {
		o.FileSettings.EnableFileAttachments = new(bool)
		*o.FileSettings.EnableFileAttachments = true
	}

	if o.FileSettings.MaxFileSize == nil {
		o.FileSettings.MaxFileSize = new(int64)
		*o.FileSettings.MaxFileSize = 52428800 // 50 MB
	}

	if o.FileSettings.PublicLinkSalt == nil || len(*o.FileSettings.PublicLinkSalt) == 0 {
		o.FileSettings.PublicLinkSalt = new(string)
		*o.FileSettings.PublicLinkSalt = NewRandomString(32)
	}

	if o.FileSettings.InitialFont == "" {
		// Defaults to "luximbi.ttf"
		o.FileSettings.InitialFont = "luximbi.ttf"
	}

	if o.FileSettings.Directory == "" {
		o.FileSettings.Directory = "./data/"
	}

	if len(o.EmailSettings.InviteSalt) == 0 {
		o.EmailSettings.InviteSalt = NewRandomString(32)
	}

	if o.PasswordSettings.MinimumLength == nil {
		o.PasswordSettings.MinimumLength = new(int)
		*o.PasswordSettings.MinimumLength = PASSWORD_MINIMUM_LENGTH
	}

	if o.PasswordSettings.Lowercase == nil {
		o.PasswordSettings.Lowercase = new(bool)
		*o.PasswordSettings.Lowercase = false
	}

	if o.PasswordSettings.Number == nil {
		o.PasswordSettings.Number = new(bool)
		*o.PasswordSettings.Number = false
	}

	if o.PasswordSettings.Uppercase == nil {
		o.PasswordSettings.Uppercase = new(bool)
		*o.PasswordSettings.Uppercase = false
	}

	if o.PasswordSettings.Symbol == nil {
		o.PasswordSettings.Symbol = new(bool)
		*o.PasswordSettings.Symbol = false
	}

	if o.EmailSettings.EnableSignInWithEmail == nil {
		o.EmailSettings.EnableSignInWithEmail = new(bool)

		if o.EmailSettings.EnableSignUpWithEmail == true {
			*o.EmailSettings.EnableSignInWithEmail = true
		} else {
			*o.EmailSettings.EnableSignInWithEmail = false
		}
	}

	if o.EmailSettings.EnableSignInWithUsername == nil {
		o.EmailSettings.EnableSignInWithUsername = new(bool)
		*o.EmailSettings.EnableSignInWithUsername = false
	}

	if o.EmailSettings.SendPushNotifications == nil {
		o.EmailSettings.SendPushNotifications = new(bool)
		*o.EmailSettings.SendPushNotifications = false
	}

	if o.EmailSettings.PushNotificationServer == nil {
		o.EmailSettings.PushNotificationServer = new(string)
		*o.EmailSettings.PushNotificationServer = ""
	}

	if o.EmailSettings.PushNotificationContents == nil {
		o.EmailSettings.PushNotificationContents = new(string)
		*o.EmailSettings.PushNotificationContents = GENERIC_NOTIFICATION
	}

	if o.EmailSettings.FeedbackOrganization == nil {
		o.EmailSettings.FeedbackOrganization = new(string)
		*o.EmailSettings.FeedbackOrganization = EMAIL_SETTINGS_DEFAULT_FEEDBACK_ORGANIZATION
	}

	if o.EmailSettings.EnableEmailBatching == nil {
		o.EmailSettings.EnableEmailBatching = new(bool)
		*o.EmailSettings.EnableEmailBatching = false
	}

	if o.EmailSettings.EmailBatchingBufferSize == nil {
		o.EmailSettings.EmailBatchingBufferSize = new(int)
		*o.EmailSettings.EmailBatchingBufferSize = EMAIL_BATCHING_BUFFER_SIZE
	}

	if o.EmailSettings.EmailBatchingInterval == nil {
		o.EmailSettings.EmailBatchingInterval = new(int)
		*o.EmailSettings.EmailBatchingInterval = EMAIL_BATCHING_INTERVAL
	}

	if o.EmailSettings.SkipServerCertificateVerification == nil {
		o.EmailSettings.SkipServerCertificateVerification = new(bool)
		*o.EmailSettings.SkipServerCertificateVerification = false
	}

	if o.LdapSettings.Enable == nil {
		o.LdapSettings.Enable = new(bool)
		*o.LdapSettings.Enable = false
	}

	if o.LdapSettings.LdapServer == nil {
		o.LdapSettings.LdapServer = new(string)
		*o.LdapSettings.LdapServer = ""
	}

	if o.LdapSettings.LdapPort == nil {
		o.LdapSettings.LdapPort = new(int)
		*o.LdapSettings.LdapPort = 389
	}

	if o.LdapSettings.ConnectionSecurity == nil {
		o.LdapSettings.ConnectionSecurity = new(string)
		*o.LdapSettings.ConnectionSecurity = ""
	}

	if o.LdapSettings.BaseDN == nil {
		o.LdapSettings.BaseDN = new(string)
		*o.LdapSettings.BaseDN = ""
	}

	if o.LdapSettings.BindUsername == nil {
		o.LdapSettings.BindUsername = new(string)
		*o.LdapSettings.BindUsername = ""
	}

	if o.LdapSettings.BindPassword == nil {
		o.LdapSettings.BindPassword = new(string)
		*o.LdapSettings.BindPassword = ""
	}

	if o.LdapSettings.UserFilter == nil {
		o.LdapSettings.UserFilter = new(string)
		*o.LdapSettings.UserFilter = ""
	}

	if o.LdapSettings.FirstNameAttribute == nil {
		o.LdapSettings.FirstNameAttribute = new(string)
		*o.LdapSettings.FirstNameAttribute = LDAP_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE
	}

	if o.LdapSettings.LastNameAttribute == nil {
		o.LdapSettings.LastNameAttribute = new(string)
		*o.LdapSettings.LastNameAttribute = LDAP_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE
	}

	if o.LdapSettings.EmailAttribute == nil {
		o.LdapSettings.EmailAttribute = new(string)
		*o.LdapSettings.EmailAttribute = LDAP_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE
	}

	if o.LdapSettings.UsernameAttribute == nil {
		o.LdapSettings.UsernameAttribute = new(string)
		*o.LdapSettings.UsernameAttribute = LDAP_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE
	}

	if o.LdapSettings.NicknameAttribute == nil {
		o.LdapSettings.NicknameAttribute = new(string)
		*o.LdapSettings.NicknameAttribute = LDAP_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE
	}

	if o.LdapSettings.IdAttribute == nil {
		o.LdapSettings.IdAttribute = new(string)
		*o.LdapSettings.IdAttribute = LDAP_SETTINGS_DEFAULT_ID_ATTRIBUTE
	}

	if o.LdapSettings.PositionAttribute == nil {
		o.LdapSettings.PositionAttribute = new(string)
		*o.LdapSettings.PositionAttribute = LDAP_SETTINGS_DEFAULT_POSITION_ATTRIBUTE
	}

	if o.LdapSettings.SyncIntervalMinutes == nil {
		o.LdapSettings.SyncIntervalMinutes = new(int)
		*o.LdapSettings.SyncIntervalMinutes = 60
	}

	if o.LdapSettings.SkipCertificateVerification == nil {
		o.LdapSettings.SkipCertificateVerification = new(bool)
		*o.LdapSettings.SkipCertificateVerification = false
	}

	if o.LdapSettings.QueryTimeout == nil {
		o.LdapSettings.QueryTimeout = new(int)
		*o.LdapSettings.QueryTimeout = 60
	}

	if o.LdapSettings.MaxPageSize == nil {
		o.LdapSettings.MaxPageSize = new(int)
		*o.LdapSettings.MaxPageSize = 0
	}

	if o.LdapSettings.LoginFieldName == nil {
		o.LdapSettings.LoginFieldName = new(string)
		*o.LdapSettings.LoginFieldName = LDAP_SETTINGS_DEFAULT_LOGIN_FIELD_NAME
	}

	if o.ClusterSettings.InterNodeListenAddress == nil {
		o.ClusterSettings.InterNodeListenAddress = new(string)
		*o.ClusterSettings.InterNodeListenAddress = ":8075"
	}

	if o.ClusterSettings.Enable == nil {
		o.ClusterSettings.Enable = new(bool)
		*o.ClusterSettings.Enable = false
	}

	if o.ClusterSettings.InterNodeUrls == nil {
		o.ClusterSettings.InterNodeUrls = []string{}
	}

	if o.MetricsSettings.ListenAddress == nil {
		o.MetricsSettings.ListenAddress = new(string)
		*o.MetricsSettings.ListenAddress = ":8067"
	}

	if o.MetricsSettings.Enable == nil {
		o.MetricsSettings.Enable = new(bool)
		*o.MetricsSettings.Enable = false
	}

	if o.AnalyticsSettings.MaxUsersForStatistics == nil {
		o.AnalyticsSettings.MaxUsersForStatistics = new(int)
		*o.AnalyticsSettings.MaxUsersForStatistics = ANALYTICS_SETTINGS_DEFAULT_MAX_USERS_FOR_STATISTICS
	}

	if o.ComplianceSettings.Enable == nil {
		o.ComplianceSettings.Enable = new(bool)
		*o.ComplianceSettings.Enable = false
	}

	if o.ComplianceSettings.Directory == nil {
		o.ComplianceSettings.Directory = new(string)
		*o.ComplianceSettings.Directory = "./data/"
	}

	if o.ComplianceSettings.EnableDaily == nil {
		o.ComplianceSettings.EnableDaily = new(bool)
		*o.ComplianceSettings.EnableDaily = false
	}

	if o.LocalizationSettings.DefaultServerLocale == nil {
		o.LocalizationSettings.DefaultServerLocale = new(string)
		*o.LocalizationSettings.DefaultServerLocale = DEFAULT_LOCALE
	}

	if o.LocalizationSettings.DefaultClientLocale == nil {
		o.LocalizationSettings.DefaultClientLocale = new(string)
		*o.LocalizationSettings.DefaultClientLocale = DEFAULT_LOCALE
	}

	if o.LocalizationSettings.AvailableLocales == nil {
		o.LocalizationSettings.AvailableLocales = new(string)
		*o.LocalizationSettings.AvailableLocales = ""
	}

	if o.SamlSettings.Enable == nil {
		o.SamlSettings.Enable = new(bool)
		*o.SamlSettings.Enable = false
	}

	if o.SamlSettings.Verify == nil {
		o.SamlSettings.Verify = new(bool)
		*o.SamlSettings.Verify = true
	}

	if o.SamlSettings.Encrypt == nil {
		o.SamlSettings.Encrypt = new(bool)
		*o.SamlSettings.Encrypt = true
	}

	if o.SamlSettings.IdpUrl == nil {
		o.SamlSettings.IdpUrl = new(string)
		*o.SamlSettings.IdpUrl = ""
	}

	if o.SamlSettings.IdpDescriptorUrl == nil {
		o.SamlSettings.IdpDescriptorUrl = new(string)
		*o.SamlSettings.IdpDescriptorUrl = ""
	}

	if o.SamlSettings.IdpCertificateFile == nil {
		o.SamlSettings.IdpCertificateFile = new(string)
		*o.SamlSettings.IdpCertificateFile = ""
	}

	if o.SamlSettings.PublicCertificateFile == nil {
		o.SamlSettings.PublicCertificateFile = new(string)
		*o.SamlSettings.PublicCertificateFile = ""
	}

	if o.SamlSettings.PrivateKeyFile == nil {
		o.SamlSettings.PrivateKeyFile = new(string)
		*o.SamlSettings.PrivateKeyFile = ""
	}

	if o.SamlSettings.AssertionConsumerServiceURL == nil {
		o.SamlSettings.AssertionConsumerServiceURL = new(string)
		*o.SamlSettings.AssertionConsumerServiceURL = ""
	}

	if o.SamlSettings.LoginButtonText == nil || *o.SamlSettings.LoginButtonText == "" {
		o.SamlSettings.LoginButtonText = new(string)
		*o.SamlSettings.LoginButtonText = USER_AUTH_SERVICE_SAML_TEXT
	}

	if o.SamlSettings.FirstNameAttribute == nil {
		o.SamlSettings.FirstNameAttribute = new(string)
		*o.SamlSettings.FirstNameAttribute = SAML_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE
	}

	if o.SamlSettings.LastNameAttribute == nil {
		o.SamlSettings.LastNameAttribute = new(string)
		*o.SamlSettings.LastNameAttribute = SAML_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE
	}

	if o.SamlSettings.EmailAttribute == nil {
		o.SamlSettings.EmailAttribute = new(string)
		*o.SamlSettings.EmailAttribute = SAML_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE
	}

	if o.SamlSettings.UsernameAttribute == nil {
		o.SamlSettings.UsernameAttribute = new(string)
		*o.SamlSettings.UsernameAttribute = SAML_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE
	}

	if o.SamlSettings.NicknameAttribute == nil {
		o.SamlSettings.NicknameAttribute = new(string)
		*o.SamlSettings.NicknameAttribute = SAML_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE
	}

	if o.SamlSettings.PositionAttribute == nil {
		o.SamlSettings.PositionAttribute = new(string)
		*o.SamlSettings.PositionAttribute = SAML_SETTINGS_DEFAULT_POSITION_ATTRIBUTE
	}

	if o.SamlSettings.LocaleAttribute == nil {
		o.SamlSettings.LocaleAttribute = new(string)
		*o.SamlSettings.LocaleAttribute = SAML_SETTINGS_DEFAULT_LOCALE_ATTRIBUTE
	}

	if o.NativeAppSettings.AppDownloadLink == nil {
		o.NativeAppSettings.AppDownloadLink = new(string)
		*o.NativeAppSettings.AppDownloadLink = NATIVEAPP_SETTINGS_DEFAULT_APP_DOWNLOAD_LINK
	}

	if o.NativeAppSettings.AndroidAppDownloadLink == nil {
		o.NativeAppSettings.AndroidAppDownloadLink = new(string)
		*o.NativeAppSettings.AndroidAppDownloadLink = NATIVEAPP_SETTINGS_DEFAULT_ANDROID_APP_DOWNLOAD_LINK
	}

	if o.NativeAppSettings.IosAppDownloadLink == nil {
		o.NativeAppSettings.IosAppDownloadLink = new(string)
		*o.NativeAppSettings.IosAppDownloadLink = NATIVEAPP_SETTINGS_DEFAULT_IOS_APP_DOWNLOAD_LINK
	}

	if o.RateLimitSettings.Enable == nil {
		o.RateLimitSettings.Enable = new(bool)
		*o.RateLimitSettings.Enable = false
	}

	if o.RateLimitSettings.MaxBurst == nil {
		o.RateLimitSettings.MaxBurst = new(int)
		*o.RateLimitSettings.MaxBurst = 100
	}

	if o.MetricsSettings.BlockProfileRate == nil {
		o.MetricsSettings.BlockProfileRate = new(int)
		*o.MetricsSettings.BlockProfileRate = 0
	}

	o.defaultWebrtcSettings()
}

func (o *Config) IsValid() *AppError {

	if *o.ClusterSettings.Enable && *o.EmailSettings.EnableEmailBatching {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.cluster_email_batching.app_error", nil, "")
	}

	if len(o.SqlSettings.AtRestEncryptKey) < 32 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.encrypt_sql.app_error", nil, "")
	}

	if !(o.SqlSettings.DriverName == DATABASE_DRIVER_MYSQL || o.SqlSettings.DriverName == DATABASE_DRIVER_POSTGRES) {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.sql_driver.app_error", nil, "")
	}

	if o.SqlSettings.MaxIdleConns <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.sql_idle.app_error", nil, "")
	}

	if len(o.SqlSettings.DataSource) == 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.sql_data_src.app_error", nil, "")
	}

	if o.SqlSettings.MaxOpenConns <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.sql_max_conn.app_error", nil, "")
	}

	if *o.FileSettings.MaxFileSize <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.max_file_size.app_error", nil, "")
	}

	if !(o.FileSettings.DriverName == IMAGE_DRIVER_LOCAL || o.FileSettings.DriverName == IMAGE_DRIVER_S3) {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_driver.app_error", nil, "")
	}

	if o.FileSettings.PreviewHeight < 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_preview_height.app_error", nil, "")
	}

	if o.FileSettings.PreviewWidth <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_preview_width.app_error", nil, "")
	}

	if o.FileSettings.ProfileHeight <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_profile_height.app_error", nil, "")
	}

	if o.FileSettings.ProfileWidth <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_profile_width.app_error", nil, "")
	}

	if o.FileSettings.ThumbnailHeight <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_thumb_height.app_error", nil, "")
	}

	if o.FileSettings.ThumbnailWidth <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_thumb_width.app_error", nil, "")
	}

	if len(*o.FileSettings.PublicLinkSalt) < 32 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.file_salt.app_error", nil, "")
	}

	if !(o.EmailSettings.ConnectionSecurity == CONN_SECURITY_NONE || o.EmailSettings.ConnectionSecurity == CONN_SECURITY_TLS || o.EmailSettings.ConnectionSecurity == CONN_SECURITY_STARTTLS || o.EmailSettings.ConnectionSecurity == CONN_SECURITY_PLAIN) {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.email_security.app_error", nil, "")
	}

	if len(o.EmailSettings.InviteSalt) < 32 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.email_salt.app_error", nil, "")
	}

	if *o.EmailSettings.EmailBatchingBufferSize <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.email_batching_buffer_size.app_error", nil, "")
	}

	if *o.EmailSettings.EmailBatchingInterval < 30 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.email_batching_interval.app_error", nil, "")
	}

	if o.RateLimitSettings.MemoryStoreSize <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.rate_mem.app_error", nil, "")
	}

	if o.RateLimitSettings.PerSec <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.rate_sec.app_error", nil, "")
	}

	if !(*o.LdapSettings.ConnectionSecurity == CONN_SECURITY_NONE || *o.LdapSettings.ConnectionSecurity == CONN_SECURITY_TLS || *o.LdapSettings.ConnectionSecurity == CONN_SECURITY_STARTTLS) {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_security.app_error", nil, "")
	}

	if *o.LdapSettings.SyncIntervalMinutes <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_sync_interval.app_error", nil, "")
	}

	if *o.LdapSettings.MaxPageSize < 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_max_page_size.app_error", nil, "")
	}

	if *o.LdapSettings.Enable {
		if *o.LdapSettings.LdapServer == "" {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_server", nil, "")
		}

		if *o.LdapSettings.BaseDN == "" {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_basedn", nil, "")
		}

		if *o.LdapSettings.EmailAttribute == "" {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_email", nil, "")
		}

		if *o.LdapSettings.UsernameAttribute == "" {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_username", nil, "")
		}

		if *o.LdapSettings.IdAttribute == "" {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_id", nil, "")
		}
	}

	if *o.SamlSettings.Enable {
		if len(*o.SamlSettings.IdpUrl) == 0 || !IsValidHttpUrl(*o.SamlSettings.IdpUrl) {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_idp_url.app_error", nil, "")
		}

		if len(*o.SamlSettings.IdpDescriptorUrl) == 0 || !IsValidHttpUrl(*o.SamlSettings.IdpDescriptorUrl) {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_idp_descriptor_url.app_error", nil, "")
		}

		if len(*o.SamlSettings.IdpCertificateFile) == 0 {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_idp_cert.app_error", nil, "")
		}

		if len(*o.SamlSettings.EmailAttribute) == 0 {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_email_attribute.app_error", nil, "")
		}

		if len(*o.SamlSettings.UsernameAttribute) == 0 {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_username_attribute.app_error", nil, "")
		}

		if *o.SamlSettings.Verify {
			if len(*o.SamlSettings.AssertionConsumerServiceURL) == 0 || !IsValidHttpUrl(*o.SamlSettings.AssertionConsumerServiceURL) {
				return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_assertion_consumer_service_url.app_error", nil, "")
			}
		}

		if *o.SamlSettings.Encrypt {
			if len(*o.SamlSettings.PrivateKeyFile) == 0 {
				return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_private_key.app_error", nil, "")
			}

			if len(*o.SamlSettings.PublicCertificateFile) == 0 {
				return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_public_cert.app_error", nil, "")
			}
		}

		if len(*o.SamlSettings.EmailAttribute) == 0 {
			return NewLocAppError("Config.IsValid", "model.config.is_valid.saml_email_attribute.app_error", nil, "")
		}
	}

	if *o.PasswordSettings.MinimumLength < PASSWORD_MINIMUM_LENGTH || *o.PasswordSettings.MinimumLength > PASSWORD_MAXIMUM_LENGTH {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.password_length.app_error", map[string]interface{}{"MinLength": PASSWORD_MINIMUM_LENGTH, "MaxLength": PASSWORD_MAXIMUM_LENGTH}, "")
	}

	if *o.RateLimitSettings.MaxBurst <= 0 {
		return NewLocAppError("Config.IsValid", "model.config.is_valid.max_burst.app_error", nil, "")
	}

	if err := o.isValidWebrtcSettings(); err != nil {
		return err
	}

	return nil
}

func (o *Config) GetSanitizeOptions() map[string]bool {
	options := map[string]bool{}
	options["fullname"] = o.PrivacySettings.ShowFullName
	options["email"] = o.PrivacySettings.ShowEmailAddress

	return options
}

func (o *Config) Sanitize() {
	if o.LdapSettings.BindPassword != nil && len(*o.LdapSettings.BindPassword) > 0 {
		*o.LdapSettings.BindPassword = FAKE_SETTING
	}

	*o.FileSettings.PublicLinkSalt = FAKE_SETTING
	if len(o.FileSettings.AmazonS3SecretAccessKey) > 0 {
		o.FileSettings.AmazonS3SecretAccessKey = FAKE_SETTING
	}

	o.EmailSettings.InviteSalt = FAKE_SETTING
	if len(o.EmailSettings.SMTPPassword) > 0 {
		o.EmailSettings.SMTPPassword = FAKE_SETTING
	}

	if len(o.GitLabSettings.Secret) > 0 {
		o.GitLabSettings.Secret = FAKE_SETTING
	}

	o.SqlSettings.DataSource = FAKE_SETTING
	o.SqlSettings.AtRestEncryptKey = FAKE_SETTING

	for i := range o.SqlSettings.DataSourceReplicas {
		o.SqlSettings.DataSourceReplicas[i] = FAKE_SETTING
	}

	for i := range o.SqlSettings.DataSourceSearchReplicas {
		o.SqlSettings.DataSourceSearchReplicas[i] = FAKE_SETTING
	}
}
