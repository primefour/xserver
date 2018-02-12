package model

const (
	SAML_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE = ""
	SAML_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE  = ""
	SAML_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE      = ""
	SAML_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE   = ""
	SAML_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE   = ""
	SAML_SETTINGS_DEFAULT_LOCALE_ATTRIBUTE     = ""
	SAML_SETTINGS_DEFAULT_POSITION_ATTRIBUTE   = ""
)

type SamlSettings struct {
	// Basic
	Enable             *bool
	EnableSyncWithLdap *bool

	Verify  *bool
	Encrypt *bool

	IdpUrl                      *string
	IdpDescriptorUrl            *string
	AssertionConsumerServiceURL *string

	IdpCertificateFile    *string
	PublicCertificateFile *string
	PrivateKeyFile        *string

	// User Mapping
	FirstNameAttribute *string
	LastNameAttribute  *string
	EmailAttribute     *string
	UsernameAttribute  *string
	NicknameAttribute  *string
	LocaleAttribute    *string
	PositionAttribute  *string

	LoginButtonText *string

	LoginButtonColor       *string
	LoginButtonBorderColor *string
	LoginButtonTextColor   *string
}

func (s *SamlSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.EnableSyncWithLdap == nil {
		s.EnableSyncWithLdap = NewBool(false)
	}

	if s.Verify == nil {
		s.Verify = NewBool(true)
	}

	if s.Encrypt == nil {
		s.Encrypt = NewBool(true)
	}

	if s.IdpUrl == nil {
		s.IdpUrl = NewString("")
	}

	if s.IdpDescriptorUrl == nil {
		s.IdpDescriptorUrl = NewString("")
	}

	if s.IdpCertificateFile == nil {
		s.IdpCertificateFile = NewString("")
	}

	if s.PublicCertificateFile == nil {
		s.PublicCertificateFile = NewString("")
	}

	if s.PrivateKeyFile == nil {
		s.PrivateKeyFile = NewString("")
	}

	if s.AssertionConsumerServiceURL == nil {
		s.AssertionConsumerServiceURL = NewString("")
	}

	if s.LoginButtonText == nil || *s.LoginButtonText == "" {
		s.LoginButtonText = NewString(USER_AUTH_SERVICE_SAML_TEXT)
	}

	if s.FirstNameAttribute == nil {
		s.FirstNameAttribute = NewString(SAML_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE)
	}

	if s.LastNameAttribute == nil {
		s.LastNameAttribute = NewString(SAML_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE)
	}

	if s.EmailAttribute == nil {
		s.EmailAttribute = NewString(SAML_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE)
	}

	if s.UsernameAttribute == nil {
		s.UsernameAttribute = NewString(SAML_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE)
	}

	if s.NicknameAttribute == nil {
		s.NicknameAttribute = NewString(SAML_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE)
	}

	if s.PositionAttribute == nil {
		s.PositionAttribute = NewString(SAML_SETTINGS_DEFAULT_POSITION_ATTRIBUTE)
	}

	if s.LocaleAttribute == nil {
		s.LocaleAttribute = NewString(SAML_SETTINGS_DEFAULT_LOCALE_ATTRIBUTE)
	}

	if s.LoginButtonColor == nil {
		s.LoginButtonColor = NewString("#34a28b")
	}

	if s.LoginButtonBorderColor == nil {
		s.LoginButtonBorderColor = NewString("#2389D7")
	}

	if s.LoginButtonTextColor == nil {
		s.LoginButtonTextColor = NewString("#ffffff")
	}
}
