package model

import (
	"github.com/primefour/xserver/utils"
)

type PrivacySettings struct {
	ShowEmailAddress bool
	ShowFullName     bool
}

func (self *PrivacySettings) setDefault() {

}

func (self *PrivacySettings) IsValidate() *utils.AppError {
	return nil
}

type SupportSettings struct {
	TermsOfServiceLink *string
	PrivacyPolicyLink  *string
	AboutLink          *string
	HelpLink           *string
	ReportAProblemLink *string
	SupportEmail       *string
}

func (self *SupportSettings) setDefault() {
	if !utils.IsSafeLink(self.TermsOfServiceLink) {
		*self.TermsOfServiceLink = SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK
	}

	if self.TermsOfServiceLink == nil {
		self.TermsOfServiceLink = new(string)
		*self.TermsOfServiceLink = SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK
	}

	if !utils.IsSafeLink(self.PrivacyPolicyLink) {
		*self.PrivacyPolicyLink = ""
	}

	if self.PrivacyPolicyLink == nil {
		self.PrivacyPolicyLink = new(string)
		*self.PrivacyPolicyLink = SUPPORT_SETTINGS_DEFAULT_PRIVACY_POLICY_LINK
	}

	if !utils.IsSafeLink(self.AboutLink) {
		*self.AboutLink = ""
	}

	if self.AboutLink == nil {
		self.AboutLink = new(string)
		*self.AboutLink = SUPPORT_SETTINGS_DEFAULT_ABOUT_LINK
	}

	if !utils.IsSafeLink(self.HelpLink) {
		*self.HelpLink = ""
	}

	if self.HelpLink == nil {
		self.HelpLink = new(string)
		*self.HelpLink = SUPPORT_SETTINGS_DEFAULT_HELP_LINK
	}

	if !utils.IsSafeLink(self.ReportAProblemLink) {
		*self.ReportAProblemLink = ""
	}

	if self.ReportAProblemLink == nil {
		self.ReportAProblemLink = new(string)
		*self.ReportAProblemLink = SUPPORT_SETTINGS_DEFAULT_REPORT_A_PROBLEM_LINK
	}

	if self.SupportEmail == nil {
		self.SupportEmail = new(string)
		*self.SupportEmail = SUPPORT_SETTINGS_DEFAULT_SUPPORT_EMAIL
	}
}

func (self *SupportSettings) IsValidate() *utils.AppError {
	return nil
}

type NativeAppSettings struct {
	AppDownloadLink        *string
	AndroidAppDownloadLink *string
	IosAppDownloadLink     *string
}

func (self *NativeAppSettings) setDefault() {
	if self.AppDownloadLink == nil {
		self.AppDownloadLink = new(string)
		*self.AppDownloadLink = NATIVEAPP_SETTINGS_DEFAULT_APP_DOWNLOAD_LINK
	}

	if self.AndroidAppDownloadLink == nil {
		self.AndroidAppDownloadLink = new(string)
		*self.AndroidAppDownloadLink = NATIVEAPP_SETTINGS_DEFAULT_ANDROID_APP_DOWNLOAD_LINK
	}

	if self.IosAppDownloadLink == nil {
		self.IosAppDownloadLink = new(string)
		*self.IosAppDownloadLink = NATIVEAPP_SETTINGS_DEFAULT_IOS_APP_DOWNLOAD_LINK
	}

}

func (self *NativeAppSettings) IsValidate() *utils.AppError {
	return nil
}
