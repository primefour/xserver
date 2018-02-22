package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	NATIVEAPP_SETTINGS_DEFAULT_APP_DOWNLOAD_LINK         = "https://about.mattermost.com/downloads/"
	NATIVEAPP_SETTINGS_DEFAULT_ANDROID_APP_DOWNLOAD_LINK = "https://about.mattermost.com/mattermost-android-app/"
	NATIVEAPP_SETTINGS_DEFAULT_IOS_APP_DOWNLOAD_LINK     = "https://about.mattermost.com/mattermost-ios-app/"

	SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK      = "https://about.mattermost.com/default-terms/"
	SUPPORT_SETTINGS_DEFAULT_PRIVACY_POLICY_LINK        = "https://about.mattermost.com/default-privacy-policy/"
	SUPPORT_SETTINGS_DEFAULT_ABOUT_LINK                 = "https://about.mattermost.com/default-about/"
	SUPPORT_SETTINGS_DEFAULT_HELP_LINK                  = "https://about.mattermost.com/default-help/"
	SUPPORT_SETTINGS_DEFAULT_REPORT_A_PROBLEM_LINK      = "https://about.mattermost.com/default-report-a-problem/"
	SUPPORT_SETTINGS_DEFAULT_ADMINISTRATORS_GUIDE_LINK  = "https://about.mattermost.com/administrators-guide/"
	SUPPORT_SETTINGS_DEFAULT_TROUBLESHOOTING_FORUM_LINK = "https://about.mattermost.com/troubleshooting-forum/"
	SUPPORT_SETTINGS_DEFAULT_COMMERCIAL_SUPPORT_LINK    = "https://about.mattermost.com/commercial-support/"
	SUPPORT_SETTINGS_DEFAULT_SUPPORT_EMAIL              = "feedback@mattermost.com"
	OTHERS_CONFIG_FILE_PATH                             = "./config/others_config.json"
	OTHERS_CONFIG_NAME                                  = "OTHERS_SETTINGS"
)

type ClientRequirements struct {
	AndroidLatestVersion string
	AndroidMinVersion    string
	DesktopLatestVersion string
	DesktopMinVersion    string
	IosLatestVersion     string
	IosMinVersion        string
}

type SupportSettings struct {
	TermsOfServiceLink *string
	PrivacyPolicyLink  *string
	AboutLink          *string
	HelpLink           *string
	ReportAProblemLink *string
	SupportEmail       *string
}

func (s *SupportSettings) SetDefaults() {
	if !IsSafeLink(s.TermsOfServiceLink) {
		*s.TermsOfServiceLink = SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK
	}

	if s.TermsOfServiceLink == nil {
		s.TermsOfServiceLink = NewString(SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK)
	}

	if !IsSafeLink(s.PrivacyPolicyLink) {
		*s.PrivacyPolicyLink = ""
	}

	if s.PrivacyPolicyLink == nil {
		s.PrivacyPolicyLink = NewString(SUPPORT_SETTINGS_DEFAULT_PRIVACY_POLICY_LINK)
	}

	if !IsSafeLink(s.AboutLink) {
		*s.AboutLink = ""
	}

	if s.AboutLink == nil {
		s.AboutLink = NewString(SUPPORT_SETTINGS_DEFAULT_ABOUT_LINK)
	}

	if !IsSafeLink(s.HelpLink) {
		*s.HelpLink = ""
	}

	if s.HelpLink == nil {
		s.HelpLink = NewString(SUPPORT_SETTINGS_DEFAULT_HELP_LINK)
	}

	if !IsSafeLink(s.ReportAProblemLink) {
		*s.ReportAProblemLink = ""
	}

	if s.ReportAProblemLink == nil {
		s.ReportAProblemLink = NewString(SUPPORT_SETTINGS_DEFAULT_REPORT_A_PROBLEM_LINK)
	}

	if s.SupportEmail == nil {
		s.SupportEmail = NewString(SUPPORT_SETTINGS_DEFAULT_SUPPORT_EMAIL)
	}
}

type ComplianceSettings struct {
	Enable      *bool
	Directory   *string
	EnableDaily *bool
}

func (s *ComplianceSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.Directory == nil {
		s.Directory = NewString("./data/")
	}

	if s.EnableDaily == nil {
		s.EnableDaily = NewBool(false)
	}
}

type NativeAppSettings struct {
	AppDownloadLink        *string
	AndroidAppDownloadLink *string
	IosAppDownloadLink     *string
}

func (s *NativeAppSettings) SetDefaults() {
	if s.AppDownloadLink == nil {
		s.AppDownloadLink = NewString(NATIVEAPP_SETTINGS_DEFAULT_APP_DOWNLOAD_LINK)
	}

	if s.AndroidAppDownloadLink == nil {
		s.AndroidAppDownloadLink = NewString(NATIVEAPP_SETTINGS_DEFAULT_ANDROID_APP_DOWNLOAD_LINK)
	}

	if s.IosAppDownloadLink == nil {
		s.IosAppDownloadLink = NewString(NATIVEAPP_SETTINGS_DEFAULT_IOS_APP_DOWNLOAD_LINK)
	}
}

type OthersSettings struct {
	ClientRequirements ClientRequirements
	NativeAppSettings  NativeAppSettings
	ComplianceSettings ComplianceSettings
	SupportSettings    SupportSettings
}

func (s *OthersSettings) SetDefaults() {
	s.NativeAppSettings.SetDefaults()
	s.ComplianceSettings.SetDefaults()
	s.SupportSettings.SetDefaults()
}

func othersConfigParser(f *os.File) (interface{}, error) {
	settings := &OthersSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("others settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetOthersSettings() *OthersSettings {
	settings := utils.GetSettings(OTHERS_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*OthersSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(OTHERS_CONFIG_NAME, OTHERS_CONFIG_FILE_PATH, true, othersConfigParser)
	if err != nil {
		return
	}
}
