package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	TEAM_SETTINGS_DEFAULT_MAX_USERS_PER_TEAM       = 50
	TEAM_SETTINGS_DEFAULT_CUSTOM_BRAND_TEXT        = ""
	TEAM_SETTINGS_DEFAULT_CUSTOM_DESCRIPTION_TEXT  = ""
	TEAM_SETTINGS_DEFAULT_USER_STATUS_AWAY_TIMEOUT = 300

	ANALYTICS_SETTINGS_DEFAULT_MAX_USERS_FOR_STATISTICS = 2500

	ANNOUNCEMENT_SETTINGS_DEFAULT_BANNER_COLOR      = "#f2a93b"
	ANNOUNCEMENT_SETTINGS_DEFAULT_BANNER_TEXT_COLOR = "#333333"

	TEAM_SETTINGS_DEFAULT_TEAM_TEXT = "default"

	DATA_RETENTION_SETTINGS_DEFAULT_MESSAGE_RETENTION_DAYS  = 365
	DATA_RETENTION_SETTINGS_DEFAULT_FILE_RETENTION_DAYS     = 365
	DATA_RETENTION_SETTINGS_DEFAULT_DELETION_JOB_START_TIME = "02:00"

	GENERIC_NO_CHANNEL_NOTIFICATION = "generic_no_channel"
	GENERIC_NOTIFICATION            = "generic"
	FULL_NOTIFICATION               = "full"

	DIRECT_MESSAGE_ANY  = "any"
	DIRECT_MESSAGE_TEAM = "team"

	SHOW_USERNAME          = "username"
	SHOW_NICKNAME_FULLNAME = "nickname_full_name"
	SHOW_FULLNAME          = "full_name"

	PERMISSIONS_ALL           = "all"
	PERMISSIONS_CHANNEL_ADMIN = "channel_admin"
	PERMISSIONS_TEAM_ADMIN    = "team_admin"
	PERMISSIONS_SYSTEM_ADMIN  = "system_admin"

	APPBASE_CONFIG_FILE_PATH = "./config/appbase_config.json"
	APPBASE_CONFIG_NAME      = "APPBASE_SETTINGS"
)

type AnalyticsSettings struct {
	MaxUsersForStatistics *int
}

func (s *AnalyticsSettings) SetDefaults() {
	if s.MaxUsersForStatistics == nil {
		s.MaxUsersForStatistics = NewInt(ANALYTICS_SETTINGS_DEFAULT_MAX_USERS_FOR_STATISTICS)
	}
}

type PasswordSettings struct {
	MinimumLength *int
	Lowercase     *bool
	Number        *bool
	Uppercase     *bool
	Symbol        *bool
}

func (s *PasswordSettings) SetDefaults() {
	if s.MinimumLength == nil {
		s.MinimumLength = NewInt(PASSWORD_MINIMUM_LENGTH)
	}

	if s.Lowercase == nil {
		s.Lowercase = NewBool(false)
	}

	if s.Number == nil {
		s.Number = NewBool(false)
	}

	if s.Uppercase == nil {
		s.Uppercase = NewBool(false)
	}

	if s.Symbol == nil {
		s.Symbol = NewBool(false)
	}
}

type RateLimitSettings struct {
	Enable           *bool
	PerSec           *int
	MaxBurst         *int
	MemoryStoreSize  *int
	VaryByRemoteAddr bool
	VaryByHeader     string
}

func (s *RateLimitSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.PerSec == nil {
		s.PerSec = NewInt(10)
	}

	if s.MaxBurst == nil {
		s.MaxBurst = NewInt(100)
	}

	if s.MemoryStoreSize == nil {
		s.MemoryStoreSize = NewInt(10000)
	}
}

func (rls *RateLimitSettings) isValid() *AppError {
	if *rls.MemoryStoreSize <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.rate_mem.app_error", nil, "", http.StatusBadRequest)
	}

	if *rls.PerSec <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.rate_sec.app_error", nil, "", http.StatusBadRequest)
	}

	if *rls.MaxBurst <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.max_burst.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

type PrivacySettings struct {
	ShowEmailAddress bool
	ShowFullName     bool
}

type AnnouncementSettings struct {
	EnableBanner         *bool
	BannerText           *string
	BannerColor          *string
	BannerTextColor      *string
	AllowBannerDismissal *bool
}

func (s *AnnouncementSettings) SetDefaults() {
	if s.EnableBanner == nil {
		s.EnableBanner = NewBool(false)
	}

	if s.BannerText == nil {
		s.BannerText = NewString("")
	}

	if s.BannerColor == nil {
		s.BannerColor = NewString(ANNOUNCEMENT_SETTINGS_DEFAULT_BANNER_COLOR)
	}

	if s.BannerTextColor == nil {
		s.BannerTextColor = NewString(ANNOUNCEMENT_SETTINGS_DEFAULT_BANNER_TEXT_COLOR)
	}

	if s.AllowBannerDismissal == nil {
		s.AllowBannerDismissal = NewBool(true)
	}
}

type ThemeSettings struct {
	EnableThemeSelection *bool
	DefaultTheme         *string
	AllowCustomThemes    *bool
	AllowedThemes        []string
}

func (s *ThemeSettings) SetDefaults() {
	if s.EnableThemeSelection == nil {
		s.EnableThemeSelection = NewBool(true)
	}

	if s.DefaultTheme == nil {
		s.DefaultTheme = NewString(TEAM_SETTINGS_DEFAULT_TEAM_TEXT)
	}

	if s.AllowCustomThemes == nil {
		s.AllowCustomThemes = NewBool(true)
	}

	if s.AllowedThemes == nil {
		s.AllowedThemes = []string{}
	}
}

type TeamSettings struct {
	SiteName                            string
	MaxUsersPerTeam                     *int
	EnableTeamCreation                  bool
	EnableUserCreation                  bool
	EnableOpenServer                    *bool
	RestrictCreationToDomains           string
	EnableCustomBrand                   *bool
	CustomBrandText                     *string
	CustomDescriptionText               *string
	RestrictDirectMessage               *string
	RestrictTeamInvite                  *string
	RestrictPublicChannelManagement     *string
	RestrictPrivateChannelManagement    *string
	RestrictPublicChannelCreation       *string
	RestrictPrivateChannelCreation      *string
	RestrictPublicChannelDeletion       *string
	RestrictPrivateChannelDeletion      *string
	RestrictPrivateChannelManageMembers *string
	EnableXToLeaveChannelsFromLHS       *bool
	UserStatusAwayTimeout               *int64
	MaxChannelsPerTeam                  *int64
	MaxNotificationsPerChannel          *int64
	EnableConfirmNotificationsToChannel *bool
	TeammateNameDisplay                 *string
	ExperimentalTownSquareIsReadOnly    *bool
	ExperimentalPrimaryTeam             *string
}

func (s *TeamSettings) SetDefaults() {
	if s.TeammateNameDisplay == nil {
		s.TeammateNameDisplay = NewString(SHOW_USERNAME)
	}
	if s.MaxUsersPerTeam == nil {
		s.MaxUsersPerTeam = NewInt(TEAM_SETTINGS_DEFAULT_MAX_USERS_PER_TEAM)
	}

	if s.EnableCustomBrand == nil {
		s.EnableCustomBrand = NewBool(false)
	}

	if s.CustomBrandText == nil {
		s.CustomBrandText = NewString(TEAM_SETTINGS_DEFAULT_CUSTOM_BRAND_TEXT)
	}

	if s.CustomDescriptionText == nil {
		s.CustomDescriptionText = NewString(TEAM_SETTINGS_DEFAULT_CUSTOM_DESCRIPTION_TEXT)
	}

	if s.EnableOpenServer == nil {
		s.EnableOpenServer = NewBool(false)
	}

	if s.RestrictDirectMessage == nil {
		s.RestrictDirectMessage = NewString(DIRECT_MESSAGE_ANY)
	}

	if s.RestrictTeamInvite == nil {
		s.RestrictTeamInvite = NewString(PERMISSIONS_ALL)
	}

	if s.RestrictPublicChannelManagement == nil {
		s.RestrictPublicChannelManagement = NewString(PERMISSIONS_ALL)
	}

	if s.RestrictPrivateChannelManagement == nil {
		s.RestrictPrivateChannelManagement = NewString(PERMISSIONS_ALL)
	}

	if s.RestrictPublicChannelCreation == nil {
		s.RestrictPublicChannelCreation = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		if *s.RestrictPublicChannelManagement == PERMISSIONS_CHANNEL_ADMIN {
			*s.RestrictPublicChannelCreation = PERMISSIONS_TEAM_ADMIN
		} else {
			*s.RestrictPublicChannelCreation = *s.RestrictPublicChannelManagement
		}
	}

	if s.RestrictPrivateChannelCreation == nil {
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		if *s.RestrictPrivateChannelManagement == PERMISSIONS_CHANNEL_ADMIN {
			s.RestrictPrivateChannelCreation = NewString(PERMISSIONS_TEAM_ADMIN)
		} else {
			s.RestrictPrivateChannelCreation = NewString(*s.RestrictPrivateChannelManagement)
		}
	}

	if s.RestrictPublicChannelDeletion == nil {
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		s.RestrictPublicChannelDeletion = NewString(*s.RestrictPublicChannelManagement)
	}

	if s.RestrictPrivateChannelDeletion == nil {
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		s.RestrictPrivateChannelDeletion = NewString(*s.RestrictPrivateChannelManagement)
	}

	if s.RestrictPrivateChannelManageMembers == nil {
		s.RestrictPrivateChannelManageMembers = NewString(PERMISSIONS_ALL)
	}

	if s.EnableXToLeaveChannelsFromLHS == nil {
		s.EnableXToLeaveChannelsFromLHS = NewBool(false)
	}

	if s.UserStatusAwayTimeout == nil {
		s.UserStatusAwayTimeout = NewInt64(TEAM_SETTINGS_DEFAULT_USER_STATUS_AWAY_TIMEOUT)
	}

	if s.MaxChannelsPerTeam == nil {
		s.MaxChannelsPerTeam = NewInt64(2000)
	}

	if s.MaxNotificationsPerChannel == nil {
		s.MaxNotificationsPerChannel = NewInt64(1000)
	}

	if s.EnableConfirmNotificationsToChannel == nil {
		s.EnableConfirmNotificationsToChannel = NewBool(true)
	}

	if s.ExperimentalTownSquareIsReadOnly == nil {
		s.ExperimentalTownSquareIsReadOnly = NewBool(false)
	}

	if s.ExperimentalPrimaryTeam == nil {
		s.ExperimentalPrimaryTeam = NewString("")
	}
}

func (ts *TeamSettings) isValid() *AppError {
	if *ts.MaxUsersPerTeam <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.max_users.app_error", nil, "", http.StatusBadRequest)
	}

	if *ts.MaxChannelsPerTeam <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.max_channels.app_error", nil, "", http.StatusBadRequest)
	}

	if *ts.MaxNotificationsPerChannel <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.max_notify_per_channel.app_error", nil, "", http.StatusBadRequest)
	}

	if !(*ts.RestrictDirectMessage == DIRECT_MESSAGE_ANY || *ts.RestrictDirectMessage == DIRECT_MESSAGE_TEAM) {
		return NewAppError("Config.IsValid", "model.config.is_valid.restrict_direct_message.app_error", nil, "", http.StatusBadRequest)
	}

	if !(*ts.TeammateNameDisplay == SHOW_FULLNAME || *ts.TeammateNameDisplay == SHOW_NICKNAME_FULLNAME || *ts.TeammateNameDisplay == SHOW_USERNAME) {
		return NewAppError("Config.IsValid", "model.config.is_valid.teammate_name_display.app_error", nil, "", http.StatusBadRequest)
	}

	if len(ts.SiteName) > SITENAME_MAX_LENGTH {
		return NewAppError("Config.IsValid", "model.config.is_valid.sitename_length.app_error", map[string]interface{}{"MaxLength": SITENAME_MAX_LENGTH}, "", http.StatusBadRequest)
	}

	return nil
}

type LocalizationSettings struct {
	DefaultServerLocale *string
	DefaultClientLocale *string
	AvailableLocales    *string
}

func (s *LocalizationSettings) SetDefaults() {
	if s.DefaultServerLocale == nil {
		s.DefaultServerLocale = NewString(DEFAULT_LOCALE)
	}

	if s.DefaultClientLocale == nil {
		s.DefaultClientLocale = NewString(DEFAULT_LOCALE)
	}

	if s.AvailableLocales == nil {
		s.AvailableLocales = NewString("")
	}
}

type DataRetentionSettings struct {
	EnableMessageDeletion *bool
	EnableFileDeletion    *bool
	MessageRetentionDays  *int
	FileRetentionDays     *int
	DeletionJobStartTime  *string
}

func (s *DataRetentionSettings) SetDefaults() {
	if s.EnableMessageDeletion == nil {
		s.EnableMessageDeletion = NewBool(false)
	}

	if s.EnableFileDeletion == nil {
		s.EnableFileDeletion = NewBool(false)
	}

	if s.MessageRetentionDays == nil {
		s.MessageRetentionDays = NewInt(DATA_RETENTION_SETTINGS_DEFAULT_MESSAGE_RETENTION_DAYS)
	}

	if s.FileRetentionDays == nil {
		s.FileRetentionDays = NewInt(DATA_RETENTION_SETTINGS_DEFAULT_FILE_RETENTION_DAYS)
	}

	if s.DeletionJobStartTime == nil {
		s.DeletionJobStartTime = NewString(DATA_RETENTION_SETTINGS_DEFAULT_DELETION_JOB_START_TIME)
	}
}

type JobSettings struct {
	RunJobs      *bool
	RunScheduler *bool
}

func (s *JobSettings) SetDefaults() {
	if s.RunJobs == nil {
		s.RunJobs = NewBool(true)
	}

	if s.RunScheduler == nil {
		s.RunScheduler = NewBool(true)
	}
}

type MessageExportSettings struct {
	EnableExport        *bool
	DailyRunTime        *string
	ExportFromTimestamp *int64
	BatchSize           *int
}

func (s *MessageExportSettings) SetDefaults() {
	if s.EnableExport == nil {
		s.EnableExport = NewBool(false)
	}

	if s.DailyRunTime == nil {
		s.DailyRunTime = NewString("01:00")
	}

	if s.ExportFromTimestamp == nil {
		s.ExportFromTimestamp = NewInt64(0)
	}

	if s.EnableExport != nil && *s.EnableExport && *s.ExportFromTimestamp == int64(0) {
		// when the feature is enabled via the System Console, use the current timestamp as the start time for future exports
		s.ExportFromTimestamp = NewInt64(GetMillis())
	} else if s.EnableExport != nil && !*s.EnableExport {
		// when the feature is disabled, reset the timestamp so that the timestamp will be set if the feature is re-enabled
		s.ExportFromTimestamp = NewInt64(0)
	}

	if s.BatchSize == nil {
		s.BatchSize = NewInt(10000)
	}
}

func (drs *DataRetentionSettings) isValid() *AppError {
	if *drs.MessageRetentionDays <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.data_retention.message_retention_days_too_low.app_error", nil, "", http.StatusBadRequest)
	}

	if *drs.FileRetentionDays <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.data_retention.file_retention_days_too_low.app_error", nil, "", http.StatusBadRequest)
	}

	if _, err := time.Parse("15:04", *drs.DeletionJobStartTime); err != nil {
		return NewAppError("Config.IsValid", "model.config.is_valid.data_retention.deletion_job_start_time.app_error", nil, err.Error(), http.StatusBadRequest)
	}

	return nil
}

func (ls *LocalizationSettings) isValid() *AppError {
	if len(*ls.AvailableLocales) > 0 {
		if !strings.Contains(*ls.AvailableLocales, *ls.DefaultClientLocale) {
			return NewAppError("Config.IsValid", "model.config.is_valid.localization.available_locales.app_error", nil, "", http.StatusBadRequest)
		}
	}

	return nil
}

func (mes *MessageExportSettings) isValid(fs FileSettings) *AppError {
	if mes.EnableExport == nil {
		return NewAppError("Config.IsValid", "model.config.is_valid.message_export.enable.app_error", nil, "", http.StatusBadRequest)
	}
	if *mes.EnableExport {
		if mes.ExportFromTimestamp == nil || *mes.ExportFromTimestamp < 0 || *mes.ExportFromTimestamp > GetMillis() {
			return NewAppError("Config.IsValid", "model.config.is_valid.message_export.export_from.app_error", nil, "", http.StatusBadRequest)
		} else if mes.DailyRunTime == nil {
			return NewAppError("Config.IsValid", "model.config.is_valid.message_export.daily_runtime.app_error", nil, "", http.StatusBadRequest)
		} else if _, err := time.Parse("15:04", *mes.DailyRunTime); err != nil {
			return NewAppError("Config.IsValid", "model.config.is_valid.message_export.daily_runtime.app_error", nil, err.Error(), http.StatusBadRequest)
		} else if mes.BatchSize == nil || *mes.BatchSize < 0 {
			return NewAppError("Config.IsValid", "model.config.is_valid.message_export.batch_size.app_error", nil, "", http.StatusBadRequest)
		}
	}
	return nil
}

type AppBaseSettings struct {
	TeamSettings          TeamSettings
	PasswordSettings      PasswordSettings
	RateLimitSettings     RateLimitSettings
	PrivacySettings       PrivacySettings
	AnnouncementSettings  AnnouncementSettings
	ThemeSettings         ThemeSettings
	LocalizationSettings  LocalizationSettings
	AnalyticsSettings     AnalyticsSettings
	DataRetentionSettings DataRetentionSettings
	MessageExportSettings MessageExportSettings
	JobSettings           JobSettings
}

func (s *AppBaseSettings) SetDefaults() {
	s.TeamSettings.SetDefaults()
	s.PasswordSettings.SetDefaults()
	s.RateLimitSettings.SetDefaults()
	s.AnnouncementSettings.SetDefaults()
	s.ThemeSettings.SetDefaults()
	s.LocalizationSettings.SetDefaults()
	s.AnalyticsSettings.SetDefaults()
	s.DataRetentionSettings.SetDefaults()
	s.MessageExportSettings.SetDefaults()
	s.JobSettings.SetDefaults()
}

func appbaseConfigParser(f *os.File) (interface{}, error) {
	settings := &AppBaseSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("AppBase settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetAppBaseSettings() *AppBaseSettings {
	settings := utils.GetSettings(APPBASE_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*AppBaseSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(APPBASE_CONFIG_NAME, APPBASE_CONFIG_FILE_PATH, true, appbaseConfigParser)
	if err != nil {
		return
	}
}
