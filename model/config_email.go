package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	EMAIL_CONFIG_FILE_PATH                       = "./config/email_config.json"
	EMAIL_CONFIG_NAME                            = "EMAIL_SETTINGS"
	EMAIL_SETTINGS_DEFAULT_FEEDBACK_ORGANIZATION = ""
	EMAIL_BATCHING_BUFFER_SIZE                   = 256
	EMAIL_BATCHING_INTERVAL                      = 30

	EMAIL_NOTIFICATION_CONTENTS_FULL    = "full"
	EMAIL_NOTIFICATION_CONTENTS_GENERIC = "generic"
)

type EmailSettings struct {
	EnableSignUpWithEmail             bool
	EnableSignInWithEmail             *bool
	EnableSignInWithUsername          *bool
	SendEmailNotifications            bool
	UseChannelInEmailNotifications    *bool
	RequireEmailVerification          bool
	FeedbackName                      string
	FeedbackEmail                     string
	FeedbackOrganization              *string
	EnableSMTPAuth                    *bool
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
	EmailNotificationContentsType     *string
	LoginButtonColor                  *string
	LoginButtonBorderColor            *string
	LoginButtonTextColor              *string
}

func (s *EmailSettings) SetDefaults() {
	if len(s.InviteSalt) == 0 {
		s.InviteSalt = NewRandomString(32)
	}

	if s.EnableSignInWithEmail == nil {
		s.EnableSignInWithEmail = NewBool(s.EnableSignUpWithEmail)
	}

	if s.EnableSignInWithUsername == nil {
		s.EnableSignInWithUsername = NewBool(false)
	}

	if s.UseChannelInEmailNotifications == nil {
		s.UseChannelInEmailNotifications = NewBool(false)
	}

	if s.SendPushNotifications == nil {
		s.SendPushNotifications = NewBool(false)
	}

	if s.PushNotificationServer == nil {
		s.PushNotificationServer = NewString("")
	}

	if s.PushNotificationContents == nil {
		s.PushNotificationContents = NewString(GENERIC_NOTIFICATION)
	}

	if s.FeedbackOrganization == nil {
		s.FeedbackOrganization = NewString(EMAIL_SETTINGS_DEFAULT_FEEDBACK_ORGANIZATION)
	}

	if s.EnableEmailBatching == nil {
		s.EnableEmailBatching = NewBool(false)
	}

	if s.EmailBatchingBufferSize == nil {
		s.EmailBatchingBufferSize = NewInt(EMAIL_BATCHING_BUFFER_SIZE)
	}

	if s.EmailBatchingInterval == nil {
		s.EmailBatchingInterval = NewInt(EMAIL_BATCHING_INTERVAL)
	}

	if s.EnableSMTPAuth == nil {
		s.EnableSMTPAuth = new(bool)
		if s.ConnectionSecurity == CONN_SECURITY_NONE {
			*s.EnableSMTPAuth = false
		} else {
			*s.EnableSMTPAuth = true
		}
	}

	if s.ConnectionSecurity == CONN_SECURITY_PLAIN {
		s.ConnectionSecurity = CONN_SECURITY_NONE
	}

	if s.SkipServerCertificateVerification == nil {
		s.SkipServerCertificateVerification = NewBool(false)
	}

	if s.EmailNotificationContentsType == nil {
		s.EmailNotificationContentsType = NewString(EMAIL_NOTIFICATION_CONTENTS_FULL)
	}

	if s.LoginButtonColor == nil {
		s.LoginButtonColor = NewString("#0000")
	}

	if s.LoginButtonBorderColor == nil {
		s.LoginButtonBorderColor = NewString("#2389D7")
	}

	if s.LoginButtonTextColor == nil {
		s.LoginButtonTextColor = NewString("#2389D7")
	}
}

func (es *EmailSettings) isValid() *AppError {
	if !(es.ConnectionSecurity == CONN_SECURITY_NONE || es.ConnectionSecurity == CONN_SECURITY_TLS || es.ConnectionSecurity == CONN_SECURITY_STARTTLS || es.ConnectionSecurity == CONN_SECURITY_PLAIN) {
		return NewAppError("Config.IsValid", "model.config.is_valid.email_security.app_error", nil, "", http.StatusBadRequest)
	}

	if len(es.InviteSalt) < 32 {
		return NewAppError("Config.IsValid", "model.config.is_valid.email_salt.app_error", nil, "", http.StatusBadRequest)
	}

	if *es.EmailBatchingBufferSize <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.email_batching_buffer_size.app_error", nil, "", http.StatusBadRequest)
	}

	if *es.EmailBatchingInterval < 30 {
		return NewAppError("Config.IsValid", "model.config.is_valid.email_batching_interval.app_error", nil, "", http.StatusBadRequest)
	}

	if !(*es.EmailNotificationContentsType == EMAIL_NOTIFICATION_CONTENTS_FULL || *es.EmailNotificationContentsType == EMAIL_NOTIFICATION_CONTENTS_GENERIC) {
		return NewAppError("Config.IsValid", "model.config.is_valid.email_notification_contents_type.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func emailConfigParser(f *os.File) (interface{}, error) {
	settings := &EmailSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("email settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetEmailSettings() *EmailSettings {
	settings := utils.GetSettings(EMAIL_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*EmailSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(EMAIL_CONFIG_NAME, EMAIL_CONFIG_FILE_PATH, true, emailConfigParser)
	if err != nil {
		return
	}
}
