package model

import (
	"encoding/json"
	"github.com/primefour/xserver/utils"
)

type SqlSettings struct {
	DriverName               string
	DataSource               string
	DataSourceReplicas       []string
	DataSourceSearchReplicas []string
	MaxIdleConns             int
	MaxOpenConns             int
	Trace                    bool
	AtRestEncryptKey         string
}

func (self *SqlSettings) setDefault() {

	if len(self.SqlSettings.AtRestEncryptKey) == 0 {
		self.SqlSettings.AtRestEncryptKey = NewRandomString(32)
	}
}

func (self *SqlSettings) IsValidate() *utils.AppError {
	if len(self.SqlSettings.AtRestEncryptKey) < 32 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.encrypt_sql.app_error", nil, "")
	}

	if !(self.SqlSettings.DriverName == DATABASE_DRIVER_MYSQL || self.SqlSettings.DriverName == DATABASE_DRIVER_POSTGRES) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.sql_driver.app_error", nil, "")
	}

	if self.SqlSettings.MaxIdleConns <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.sql_idle.app_error", nil, "")
	}

	if len(self.SqlSettings.DataSource) == 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.sql_data_src.app_error", nil, "")
	}

	if self.SqlSettings.MaxOpenConns <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.sql_max_conn.app_error", nil, "")
	}
}

type PasswordSettings struct {
	MinimumLength *int
	Lowercase     *bool
	Number        *bool
	Uppercase     *bool
	Symbol        *bool
}

func (self *PasswordSettings) setDefault() {
	if self.PasswordSettings.MinimumLength == nil {
		self.PasswordSettings.MinimumLength = new(int)
		*self.PasswordSettings.MinimumLength = PASSWORD_MINIMUM_LENGTH
	}

	if self.PasswordSettings.Lowercase == nil {
		self.PasswordSettings.Lowercase = new(bool)
		*self.PasswordSettings.Lowercase = false
	}

	if self.PasswordSettings.Number == nil {
		self.PasswordSettings.Number = new(bool)
		*self.PasswordSettings.Number = false
	}

	if self.PasswordSettings.Uppercase == nil {
		self.PasswordSettings.Uppercase = new(bool)
		*self.PasswordSettings.Uppercase = false
	}

	if self.PasswordSettings.Symbol == nil {
		self.PasswordSettings.Symbol = new(bool)
		*self.PasswordSettings.Symbol = false
	}
}

func (self *PasswordSettings) IsValidate() utils.AppError {
	if *self.PasswordSettings.MinimumLength < PASSWORD_MINIMUM_LENGTH || *self.PasswordSettings.MinimumLength > PASSWORD_MAXIMUM_LENGTH {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.password_length.app_error", map[string]interface{}{"MinLength": PASSWORD_MINIMUM_LENGTH, "MaxLength": PASSWORD_MAXIMUM_LENGTH}, "")
	}
}

type FileSettings struct {
	EnableFileAttachments   *bool
	MaxFileSize             *int64
	DriverName              string
	Directory               string
	EnablePublicLink        bool
	PublicLinkSalt          *string
	ThumbnailWidth          int
	ThumbnailHeight         int
	PreviewWidth            int
	PreviewHeight           int
	ProfileWidth            int
	ProfileHeight           int
	InitialFont             string
	AmazonS3AccessKeyId     string
	AmazonS3SecretAccessKey string
	AmazonS3Bucket          string
	AmazonS3Region          string
	AmazonS3Endpoint        string
	AmazonS3SSL             *bool
}

func (self *FileSettings) setDefault() {
	if self.AmazonS3Endpoint == "" {
		// Defaults to "s3.amazonaws.com"
		self.AmazonS3Endpoint = "s3.amazonaws.com"
	}

	if self.AmazonS3Region == "" {
		// Defaults to "us-east-1" region.
		self.AmazonS3Region = "us-east-1"
	}

	if self.AmazonS3SSL == nil {
		self.AmazonS3SSL = new(bool)
		*self.AmazonS3SSL = true // Secure by default.
	}

	if self.EnableFileAttachments == nil {
		self.EnableFileAttachments = new(bool)
		*self.EnableFileAttachments = true
	}

	if self.MaxFileSize == nil {
		self.MaxFileSize = new(int64)
		*self.MaxFileSize = 52428800 // 50 MB
	}

	if self.PublicLinkSalt == nil || len(*self.PublicLinkSalt) == 0 {
		self.PublicLinkSalt = new(string)
		*self.PublicLinkSalt = utils.NewRandomString(32)
	}

	if self.InitialFont == "" {
		// Defaults to "luximbi.ttf"
		self.InitialFont = "luximbi.ttf"
	}

	if self.Directory == "" {
		self.Directory = "./data/"
	}
}

func (self *FileSettings) IsValidate() *utils.AppError {

	if *self.MaxFileSize <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_file_size.app_error", nil, "")
	}

	if !(self.DriverName == IMAGE_DRIVER_LOCAL || self.DriverName == IMAGE_DRIVER_S3) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_driver.app_error", nil, "")
	}

	if self.PreviewHeight < 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_preview_height.app_error", nil, "")
	}

	if self.PreviewWidth <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_preview_width.app_error", nil, "")
	}

	if self.ProfileHeight <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_profile_height.app_error", nil, "")
	}

	if self.ProfileWidth <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_profile_width.app_error", nil, "")
	}

	if self.ThumbnailHeight <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_thumb_height.app_error", nil, "")
	}

	if self.ThumbnailWidth <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_thumb_width.app_error", nil, "")
	}

	if len(*self.PublicLinkSalt) < 32 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.file_salt.app_error", nil, "")
	}
}

type EmailSettings struct {
	EnableSignUpWithEmail             bool
	EnableSignInWithEmail             *bool
	EnableSignInWithUsername          *bool
	SendEmailNotifications            bool
	RequireEmailVerification          bool
	FeedbackName                      string
	FeedbackEmail                     string
	FeedbackOrganization              *string
	SMTPUsername                      string
	SMTPPassword                      string
	SMTPServer                        string
	SMTPPort                          string
	ConnectionSecurity                string
	InviteSalt                        string
	SendPushNotifications             *bool
	PushNotificationServer            *string
	PushNotificationContents          *string
	EnableEmailBatching               *bool
	EmailBatchingBufferSize           *int
	EmailBatchingInterval             *int
	SkipServerCertificateVerification *bool
}

func (self *EmailSettings) setDefault() {
	if len(self.InviteSalt) == 0 {
		self.InviteSalt = NewRandomString(32)
	}

	if self.EnableSignInWithEmail == nil {
		self.EnableSignInWithEmail = new(bool)

		if self.EnableSignUpWithEmail == true {
			*self.EnableSignInWithEmail = true
		} else {
			*self.EnableSignInWithEmail = false
		}
	}

	if self.EnableSignInWithUsername == nil {
		self.EnableSignInWithUsername = new(bool)
		*self.EnableSignInWithUsername = false
	}

	if self.SendPushNotifications == nil {
		self.SendPushNotifications = new(bool)
		*self.SendPushNotifications = false
	}

	if self.PushNotificationServer == nil {
		self.PushNotificationServer = new(string)
		*self.PushNotificationServer = ""
	}

	if self.PushNotificationContents == nil {
		self.PushNotificationContents = new(string)
		*self.PushNotificationContents = GENERIC_NOTIFICATION
	}

	if self.FeedbackOrganization == nil {
		self.FeedbackOrganization = new(string)
		*self.FeedbackOrganization = EMAIL_SETTINGS_DEFAULT_FEEDBACK_ORGANIZATION
	}

	if self.EnableEmailBatching == nil {
		self.EnableEmailBatching = new(bool)
		*self.EnableEmailBatching = false
	}

	if self.EmailBatchingBufferSize == nil {
		self.EmailBatchingBufferSize = new(int)
		*self.EmailBatchingBufferSize = EMAIL_BATCHING_BUFFER_SIZE
	}

	if self.EmailBatchingInterval == nil {
		self.EmailBatchingInterval = new(int)
		*self.EmailBatchingInterval = EMAIL_BATCHING_INTERVAL
	}

	if self.SkipServerCertificateVerification == nil {
		self.SkipServerCertificateVerification = new(bool)
		*self.SkipServerCertificateVerification = false
	}
}

func (self *EmailSettings) IsValidate() *utils.AppError {

	if !(self.ConnectionSecurity == CONN_SECURITY_NONE || self.ConnectionSecurity == CONN_SECURITY_TLS || self.ConnectionSecurity == CONN_SECURITY_STARTTLS || self.ConnectionSecurity == CONN_SECURITY_PLAIN) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.email_security.app_error", nil, "")
	}

	if len(self.InviteSalt) < 32 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.email_salt.app_error", nil, "")
	}

	if *self.EmailBatchingBufferSize <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.email_batching_buffer_size.app_error", nil, "")
	}

	if *self.EmailBatchingInterval < 30 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.email_batching_interval.app_error", nil, "")
	}
}

type RateLimitSettings struct {
	Enable           *bool
	PerSec           int
	MaxBurst         *int
	MemoryStoreSize  int
	VaryByRemoteAddr bool
	VaryByHeader     string
}

func (self *RateLimitSettings) setDefault() {
	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}
	if self.MaxBurst == nil {
		self.MaxBurst = new(int)
		*self.MaxBurst = 100
	}
}

func (self *RateLimitSettings) IsValidate() *utils.AppError {
	if self.MemoryStoreSize <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.rate_mem.app_error", nil, "")
	}

	if self.PerSec <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.rate_sec.app_error", nil, "")
	}

	if *self.MaxBurst <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_burst.app_error", nil, "")
	}
}
