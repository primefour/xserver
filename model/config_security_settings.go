package model

type SamlSettings struct {
	// Basic
	Enable  *bool
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
}

func (self *SamlSettings) setDefault() {

	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}

	if self.Verify == nil {
		self.Verify = new(bool)
		*self.Verify = true
	}

	if self.Encrypt == nil {
		self.Encrypt = new(bool)
		*self.Encrypt = true
	}

	if self.IdpUrl == nil {
		self.IdpUrl = new(string)
		*self.IdpUrl = ""
	}

	if self.IdpDescriptorUrl == nil {
		self.IdpDescriptorUrl = new(string)
		*self.IdpDescriptorUrl = ""
	}

	if self.IdpCertificateFile == nil {
		self.IdpCertificateFile = new(string)
		*self.IdpCertificateFile = ""
	}

	if self.PublicCertificateFile == nil {
		self.PublicCertificateFile = new(string)
		*self.PublicCertificateFile = ""
	}

	if self.PrivateKeyFile == nil {
		self.PrivateKeyFile = new(string)
		*self.PrivateKeyFile = ""
	}

	if self.AssertionConsumerServiceURL == nil {
		self.AssertionConsumerServiceURL = new(string)
		*self.AssertionConsumerServiceURL = ""
	}

	if self.LoginButtonText == nil || *self.LoginButtonText == "" {
		self.LoginButtonText = new(string)
		*self.LoginButtonText = USER_AUTH_SERVICE_SAML_TEXT
	}

	if self.FirstNameAttribute == nil {
		self.FirstNameAttribute = new(string)
		*self.FirstNameAttribute = SAML_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE
	}

	if self.LastNameAttribute == nil {
		self.LastNameAttribute = new(string)
		*self.LastNameAttribute = SAML_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE
	}

	if self.EmailAttribute == nil {
		self.EmailAttribute = new(string)
		*self.EmailAttribute = SAML_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE
	}

	if self.UsernameAttribute == nil {
		self.UsernameAttribute = new(string)
		*self.UsernameAttribute = SAML_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE
	}

	if self.NicknameAttribute == nil {
		self.NicknameAttribute = new(string)
		*self.NicknameAttribute = SAML_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE
	}

	if self.PositionAttribute == nil {
		self.PositionAttribute = new(string)
		*self.PositionAttribute = SAML_SETTINGS_DEFAULT_POSITION_ATTRIBUTE
	}

	if self.LocaleAttribute == nil {
		self.LocaleAttribute = new(string)
		*self.LocaleAttribute = SAML_SETTINGS_DEFAULT_LOCALE_ATTRIBUTE
	}
}

func (self *SamlSettings) IsValidate() *utils.AppError {

	if *self.Enable {
		if len(*self.IdpUrl) == 0 || !IsValidHttpUrl(*self.IdpUrl) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_idp_url.app_error", nil, "")
		}

		if len(*self.IdpDescriptorUrl) == 0 || !IsValidHttpUrl(*self.IdpDescriptorUrl) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_idp_descriptor_url.app_error", nil, "")
		}

		if len(*self.IdpCertificateFile) == 0 {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_idp_cert.app_error", nil, "")
		}

		if len(*self.EmailAttribute) == 0 {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_email_attribute.app_error", nil, "")
		}

		if len(*self.UsernameAttribute) == 0 {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_username_attribute.app_error", nil, "")
		}

		if *self.Verify {
			if len(*self.AssertionConsumerServiceURL) == 0 || !IsValidHttpUrl(*self.AssertionConsumerServiceURL) {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_assertion_consumer_service_url.app_error", nil, "")
			}
		}

		if *self.Encrypt {
			if len(*self.PrivateKeyFile) == 0 {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_private_key.app_error", nil, "")
			}

			if len(*self.PublicCertificateFile) == 0 {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_public_cert.app_error", nil, "")
			}
		}

		if len(*self.EmailAttribute) == 0 {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.saml_email_attribute.app_error", nil, "")
		}
	}
	return nil
}

type LdapSettings struct {
	// Basic
	Enable             *bool
	LdapServer         *string
	LdapPort           *int
	ConnectionSecurity *string
	BaseDN             *string
	BindUsername       *string
	BindPassword       *string

	// Filtering
	UserFilter *string

	// User Mapping
	FirstNameAttribute *string
	LastNameAttribute  *string
	EmailAttribute     *string
	UsernameAttribute  *string
	NicknameAttribute  *string
	IdAttribute        *string
	PositionAttribute  *string

	// Syncronization
	SyncIntervalMinutes *int

	// Advanced
	SkipCertificateVerification *bool
	QueryTimeout                *int
	MaxPageSize                 *int

	// Customization
	LoginFieldName *string
}

func (self *LdapSettings) setDefault() {

	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}

	if self.LdapServer == nil {
		self.LdapServer = new(string)
		*self.LdapServer = ""
	}

	if self.LdapPort == nil {
		self.LdapPort = new(int)
		*self.LdapPort = 389
	}

	if self.ConnectionSecurity == nil {
		self.ConnectionSecurity = new(string)
		*self.ConnectionSecurity = ""
	}

	if self.BaseDN == nil {
		self.BaseDN = new(string)
		*self.BaseDN = ""
	}

	if self.BindUsername == nil {
		self.BindUsername = new(string)
		*self.BindUsername = ""
	}

	if self.BindPassword == nil {
		self.BindPassword = new(string)
		*self.BindPassword = ""
	}

	if self.UserFilter == nil {
		self.UserFilter = new(string)
		*self.UserFilter = ""
	}

	if self.FirstNameAttribute == nil {
		self.FirstNameAttribute = new(string)
		*self.FirstNameAttribute = LDAP_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE
	}

	if self.LastNameAttribute == nil {
		self.LastNameAttribute = new(string)
		*self.LastNameAttribute = LDAP_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE
	}

	if self.EmailAttribute == nil {
		self.EmailAttribute = new(string)
		*self.EmailAttribute = LDAP_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE
	}

	if self.UsernameAttribute == nil {
		self.UsernameAttribute = new(string)
		*self.UsernameAttribute = LDAP_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE
	}

	if self.NicknameAttribute == nil {
		self.NicknameAttribute = new(string)
		*self.NicknameAttribute = LDAP_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE
	}

	if self.IdAttribute == nil {
		self.IdAttribute = new(string)
		*self.IdAttribute = LDAP_SETTINGS_DEFAULT_ID_ATTRIBUTE
	}

	if self.PositionAttribute == nil {
		self.PositionAttribute = new(string)
		*self.PositionAttribute = LDAP_SETTINGS_DEFAULT_POSITION_ATTRIBUTE
	}

	if self.SyncIntervalMinutes == nil {
		self.SyncIntervalMinutes = new(int)
		*self.SyncIntervalMinutes = 60
	}

	if self.SkipCertificateVerification == nil {
		self.SkipCertificateVerification = new(bool)
		*self.SkipCertificateVerification = false
	}

	if self.QueryTimeout == nil {
		self.QueryTimeout = new(int)
		*self.QueryTimeout = 60
	}

	if self.MaxPageSize == nil {
		self.MaxPageSize = new(int)
		*self.MaxPageSize = 0
	}

	if self.LoginFieldName == nil {
		self.LoginFieldName = new(string)
		*self.LoginFieldName = LDAP_SETTINGS_DEFAULT_LOGIN_FIELD_NAME
	}
}

func (self *LdapSettings) IsValidate() *utils.AppError {

	if !(*self.ConnectionSecurity == CONN_SECURITY_NONE || *self.ConnectionSecurity == CONN_SECURITY_TLS || *self.ConnectionSecurity == CONN_SECURITY_STARTTLS) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_security.app_error", nil, "")
	}

	if *self.SyncIntervalMinutes <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_sync_interval.app_error", nil, "")
	}

	if *self.MaxPageSize < 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_max_page_size.app_error", nil, "")
	}

	if *self.Enable {
		if *self.LdapServer == "" {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_server", nil, "")
		}

		if *self.BaseDN == "" {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_basedn", nil, "")
		}

		if *self.EmailAttribute == "" {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_email", nil, "")
		}

		if *self.UsernameAttribute == "" {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_username", nil, "")
		}

		if *self.IdAttribute == "" {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.ldap_id", nil, "")
		}
	}

}

type SSOSettings struct {
	Enable          bool
	Secret          string
	Id              string
	Scope           string
	AuthEndpoint    string
	TokenEndpoint   string
	UserApiEndpoint string
}

func (self *SSOSettings) setDefault() {
}

func (self *SSOSettings) IsValidate() *utils.AppError {

}
