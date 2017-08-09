package model

import (
	"github.com/primefour/xserver/utils"
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

func (self *ClusterSettings) IsValidate() *utils.AppError {
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

func (self *AnalyticsSettings) IsValidate() *utils.AppError {
	return nil
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
