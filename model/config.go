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

func (self *ClusterSettings) setDefault() {
	if self.InterNodeListenAddress == nil {
		self.InterNodeListenAddress = new(string)
		*self.InterNodeListenAddress = ":8075"
	}

	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}

	if self.InterNodeUrls == nil {
		self.InterNodeUrls = []string{}
	}

}

func (self *ClusterSettings) IsValidate() utils.AppError {
	return nil
}

type MetricsSettings struct {
	Enable           *bool
	BlockProfileRate *int
	ListenAddress    *string
}

func (self *MetricsSettings) setDefault() {

	if self.ListenAddress == nil {
		self.ListenAddress = new(string)
		*self.ListenAddress = ":8067"
	}

	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}

	if self.BlockProfileRate == nil {
		self.BlockProfileRate = new(int)
		*self.BlockProfileRate = 0
	}
}

func (self *MetricsSettings) IsValidate() *utils.AppError {
	return nil
}

type AnalyticsSettings struct {
	MaxUsersForStatistics *int
}

func (self *AnalyticsSettings) setDefault() {
	if self.MaxUsersForStatistics == nil {
		self.MaxUsersForStatistics = new(int)
		*self.MaxUsersForStatistics = ANALYTICS_SETTINGS_DEFAULT_MAX_USERS_FOR_STATISTICS
	}
}

func (self *AnalyticsSettings) IsValidate() utils.AppError {

}

type ComplianceSettings struct {
	Enable      *bool
	Directory   *string
	EnableDaily *bool
}

func (self *ComplianceSettings) setDefault() {

	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}

	if self.Directory == nil {
		self.Directory = new(string)
		*self.Directory = "./data/"
	}

	if self.EnableDaily == nil {
		self.EnableDaily = new(bool)
		*self.EnableDaily = false
	}
}

func (self *ComplianceSettings) IsValidate() *utils.AppError {
	return nil
}

type LocalizationSettings struct {
	DefaultServerLocale *string
	DefaultClientLocale *string
	AvailableLocales    *string
}

func (self *LocalizationSettings) setDefault() {

	if self.DefaultServerLocale == nil {
		self.DefaultServerLocale = new(string)
		*self.DefaultServerLocale = utils.DEFAULT_LOCALE
	}

	if self.DefaultClientLocale == nil {
		self.DefaultClientLocale = new(string)
		*self.DefaultClientLocale = utils.DEFAULT_LOCALE
	}

	if self.AvailableLocales == nil {
		self.AvailableLocales = new(string)
		*self.AvailableLocales = ""
	}
}

func (self *LocalizationSettings) IsValidate() *utils.AppError {
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
