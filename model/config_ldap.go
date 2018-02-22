package model

const (
	LDAP_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE = ""
	LDAP_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE  = ""
	LDAP_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE      = ""
	LDAP_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE   = ""
	LDAP_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE   = ""
	LDAP_SETTINGS_DEFAULT_ID_ATTRIBUTE         = ""
	LDAP_SETTINGS_DEFAULT_POSITION_ATTRIBUTE   = ""
	LDAP_SETTINGS_DEFAULT_LOGIN_FIELD_NAME     = ""
)

type LdapSettings struct {
	// Basic
	Enable             *bool
	EnableSync         *bool
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

	LoginButtonColor       *string
	LoginButtonBorderColor *string
	LoginButtonTextColor   *string
}

func (s *LdapSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	// When unset should default to LDAP Enabled
	if s.EnableSync == nil {
		s.EnableSync = NewBool(*s.Enable)
	}

	if s.LdapServer == nil {
		s.LdapServer = NewString("")
	}

	if s.LdapPort == nil {
		s.LdapPort = NewInt(389)
	}

	if s.ConnectionSecurity == nil {
		s.ConnectionSecurity = NewString("")
	}

	if s.BaseDN == nil {
		s.BaseDN = NewString("")
	}

	if s.BindUsername == nil {
		s.BindUsername = NewString("")
	}

	if s.BindPassword == nil {
		s.BindPassword = NewString("")
	}

	if s.UserFilter == nil {
		s.UserFilter = NewString("")
	}

	if s.FirstNameAttribute == nil {
		s.FirstNameAttribute = NewString(LDAP_SETTINGS_DEFAULT_FIRST_NAME_ATTRIBUTE)
	}

	if s.LastNameAttribute == nil {
		s.LastNameAttribute = NewString(LDAP_SETTINGS_DEFAULT_LAST_NAME_ATTRIBUTE)
	}

	if s.EmailAttribute == nil {
		s.EmailAttribute = NewString(LDAP_SETTINGS_DEFAULT_EMAIL_ATTRIBUTE)
	}

	if s.UsernameAttribute == nil {
		s.UsernameAttribute = NewString(LDAP_SETTINGS_DEFAULT_USERNAME_ATTRIBUTE)
	}

	if s.NicknameAttribute == nil {
		s.NicknameAttribute = NewString(LDAP_SETTINGS_DEFAULT_NICKNAME_ATTRIBUTE)
	}

	if s.IdAttribute == nil {
		s.IdAttribute = NewString(LDAP_SETTINGS_DEFAULT_ID_ATTRIBUTE)
	}

	if s.PositionAttribute == nil {
		s.PositionAttribute = NewString(LDAP_SETTINGS_DEFAULT_POSITION_ATTRIBUTE)
	}

	if s.SyncIntervalMinutes == nil {
		s.SyncIntervalMinutes = NewInt(60)
	}

	if s.SkipCertificateVerification == nil {
		s.SkipCertificateVerification = NewBool(false)
	}

	if s.QueryTimeout == nil {
		s.QueryTimeout = NewInt(60)
	}

	if s.MaxPageSize == nil {
		s.MaxPageSize = NewInt(0)
	}

	if s.LoginFieldName == nil {
		s.LoginFieldName = NewString(LDAP_SETTINGS_DEFAULT_LOGIN_FIELD_NAME)
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

func (ls *LdapSettings) isValid() *AppError {
	if !(*ls.ConnectionSecurity == CONN_SECURITY_NONE || *ls.ConnectionSecurity == CONN_SECURITY_TLS || *ls.ConnectionSecurity == CONN_SECURITY_STARTTLS) {
		return NewAppError("Config.IsValid", "model.config.is_valid.ldap_security.app_error", nil, "", http.StatusBadRequest)
	}

	if *ls.SyncIntervalMinutes <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.ldap_sync_interval.app_error", nil, "", http.StatusBadRequest)
	}

	if *ls.MaxPageSize < 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.ldap_max_page_size.app_error", nil, "", http.StatusBadRequest)
	}

	if *ls.Enable {
		if *ls.LdapServer == "" {
			return NewAppError("Config.IsValid", "model.config.is_valid.ldap_server", nil, "", http.StatusBadRequest)
		}

		if *ls.BaseDN == "" {
			return NewAppError("Config.IsValid", "model.config.is_valid.ldap_basedn", nil, "", http.StatusBadRequest)
		}

		if *ls.EmailAttribute == "" {
			return NewAppError("Config.IsValid", "model.config.is_valid.ldap_email", nil, "", http.StatusBadRequest)
		}

		if *ls.UsernameAttribute == "" {
			return NewAppError("Config.IsValid", "model.config.is_valid.ldap_username", nil, "", http.StatusBadRequest)
		}

		if *ls.IdAttribute == "" {
			return NewAppError("Config.IsValid", "model.config.is_valid.ldap_id", nil, "", http.StatusBadRequest)
		}
	}

	return nil
}
