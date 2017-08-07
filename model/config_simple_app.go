package model

import (
	"encoding/json"
	"github.com/primefour/xserver/utils"
)

type PrivacySettings struct {
	ShowEmailAddress bool
	ShowFullName     bool
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
	if !utils.IsSafeLink(self.SupportSettings.TermsOfServiceLink) {
		*self.SupportSettings.TermsOfServiceLink = SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK
	}

	if self.SupportSettings.TermsOfServiceLink == nil {
		self.SupportSettings.TermsOfServiceLink = new(string)
		*self.SupportSettings.TermsOfServiceLink = SUPPORT_SETTINGS_DEFAULT_TERMS_OF_SERVICE_LINK
	}

	if !utils.IsSafeLink(self.SupportSettings.PrivacyPolicyLink) {
		*self.SupportSettings.PrivacyPolicyLink = ""
	}

	if self.SupportSettings.PrivacyPolicyLink == nil {
		self.SupportSettings.PrivacyPolicyLink = new(string)
		*self.SupportSettings.PrivacyPolicyLink = SUPPORT_SETTINGS_DEFAULT_PRIVACY_POLICY_LINK
	}

	if !utils.IsSafeLink(self.SupportSettings.AboutLink) {
		*self.SupportSettings.AboutLink = ""
	}

	if self.SupportSettings.AboutLink == nil {
		self.SupportSettings.AboutLink = new(string)
		*self.SupportSettings.AboutLink = SUPPORT_SETTINGS_DEFAULT_ABOUT_LINK
	}

	if !utils.IsSafeLink(self.SupportSettings.HelpLink) {
		*self.SupportSettings.HelpLink = ""
	}

	if self.SupportSettings.HelpLink == nil {
		self.SupportSettings.HelpLink = new(string)
		*self.SupportSettings.HelpLink = SUPPORT_SETTINGS_DEFAULT_HELP_LINK
	}

	if !utils.IsSafeLink(self.SupportSettings.ReportAProblemLink) {
		*self.SupportSettings.ReportAProblemLink = ""
	}

	if self.SupportSettings.ReportAProblemLink == nil {
		self.SupportSettings.ReportAProblemLink = new(string)
		*self.SupportSettings.ReportAProblemLink = SUPPORT_SETTINGS_DEFAULT_REPORT_A_PROBLEM_LINK
	}

	if self.SupportSettings.SupportEmail == nil {
		self.SupportSettings.SupportEmail = new(string)
		*self.SupportSettings.SupportEmail = SUPPORT_SETTINGS_DEFAULT_SUPPORT_EMAIL
	}
}

type TeamSettings struct {
	SiteName                            string
	MaxUsersPerTeam                     int
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
	UserStatusAwayTimeout               *int64
	MaxChannelsPerTeam                  *int64
	MaxNotificationsPerChannel          *int64
}

func (self *TeamSettings) setDefault() {
	if self.TeamSettings.EnableCustomBrand == nil {
		self.TeamSettings.EnableCustomBrand = new(bool)
		*self.TeamSettings.EnableCustomBrand = false
	}

	if self.TeamSettings.CustomBrandText == nil {
		self.TeamSettings.CustomBrandText = new(string)
		*self.TeamSettings.CustomBrandText = TEAM_SETTINGS_DEFAULT_CUSTOM_BRAND_TEXT
	}

	if self.TeamSettings.CustomDescriptionText == nil {
		self.TeamSettings.CustomDescriptionText = new(string)
		*self.TeamSettings.CustomDescriptionText = TEAM_SETTINGS_DEFAULT_CUSTOM_DESCRIPTION_TEXT
	}

	if self.TeamSettings.EnableOpenServer == nil {
		self.TeamSettings.EnableOpenServer = new(bool)
		*self.TeamSettings.EnableOpenServer = false
	}

	if self.TeamSettings.RestrictDirectMessage == nil {
		self.TeamSettings.RestrictDirectMessage = new(string)
		*self.TeamSettings.RestrictDirectMessage = DIRECT_MESSAGE_ANY
	}

	if self.TeamSettings.RestrictTeamInvite == nil {
		self.TeamSettings.RestrictTeamInvite = new(string)
		*self.TeamSettings.RestrictTeamInvite = PERMISSIONS_ALL
	}

	if self.TeamSettings.RestrictPublicChannelManagement == nil {
		self.TeamSettings.RestrictPublicChannelManagement = new(string)
		*self.TeamSettings.RestrictPublicChannelManagement = PERMISSIONS_ALL
	}

	if self.TeamSettings.RestrictPrivateChannelManagement == nil {
		self.TeamSettings.RestrictPrivateChannelManagement = new(string)
		*self.TeamSettings.RestrictPrivateChannelManagement = PERMISSIONS_ALL
	}

	if self.TeamSettings.RestrictPublicChannelCreation == nil {
		self.TeamSettings.RestrictPublicChannelCreation = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		if *self.TeamSettings.RestrictPublicChannelManagement == PERMISSIONS_CHANNEL_ADMIN {
			*self.TeamSettings.RestrictPublicChannelCreation = PERMISSIONS_TEAM_ADMIN
		} else {
			*self.TeamSettings.RestrictPublicChannelCreation = *self.TeamSettings.RestrictPublicChannelManagement
		}
	}

	if self.TeamSettings.RestrictPrivateChannelCreation == nil {
		self.TeamSettings.RestrictPrivateChannelCreation = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		if *self.TeamSettings.RestrictPrivateChannelManagement == PERMISSIONS_CHANNEL_ADMIN {
			*self.TeamSettings.RestrictPrivateChannelCreation = PERMISSIONS_TEAM_ADMIN
		} else {
			*self.TeamSettings.RestrictPrivateChannelCreation = *self.TeamSettings.RestrictPrivateChannelManagement
		}
	}

	if self.TeamSettings.RestrictPublicChannelDeletion == nil {
		self.TeamSettings.RestrictPublicChannelDeletion = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		*self.TeamSettings.RestrictPublicChannelDeletion = *self.TeamSettings.RestrictPublicChannelManagement
	}

	if self.TeamSettings.RestrictPrivateChannelDeletion == nil {
		self.TeamSettings.RestrictPrivateChannelDeletion = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		*self.TeamSettings.RestrictPrivateChannelDeletion = *self.TeamSettings.RestrictPrivateChannelManagement
	}

	if self.TeamSettings.RestrictPrivateChannelManageMembers == nil {
		self.TeamSettings.RestrictPrivateChannelManageMembers = new(string)
		*self.TeamSettings.RestrictPrivateChannelManageMembers = PERMISSIONS_ALL
	}

	if self.TeamSettings.UserStatusAwayTimeout == nil {
		self.TeamSettings.UserStatusAwayTimeout = new(int64)
		*self.TeamSettings.UserStatusAwayTimeout = TEAM_SETTINGS_DEFAULT_USER_STATUS_AWAY_TIMEOUT
	}

	if self.TeamSettings.MaxChannelsPerTeam == nil {
		self.TeamSettings.MaxChannelsPerTeam = new(int64)
		*self.TeamSettings.MaxChannelsPerTeam = 2000
	}

	if self.TeamSettings.MaxNotificationsPerChannel == nil {
		self.TeamSettings.MaxNotificationsPerChannel = new(int64)
		*self.TeamSettings.MaxNotificationsPerChannel = 1000
	}
}

func (self *TeamSettings) IsValidate() *utils.AppError {

	if self.TeamSettings.MaxUsersPerTeam <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_users.app_error", nil, "")
	}

	if *self.TeamSettings.MaxChannelsPerTeam <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_channels.app_error", nil, "")
	}

	if *self.TeamSettings.MaxNotificationsPerChannel <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_notify_per_channel.app_error", nil, "")
	}

	if !(*self.TeamSettings.RestrictDirectMessage == DIRECT_MESSAGE_ANY || *self.TeamSettings.RestrictDirectMessage == DIRECT_MESSAGE_TEAM) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.restrict_direct_message.app_error", nil, "")
	}

	if len(self.TeamSettings.SiteName) > SITENAME_MAX_LENGTH {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.sitename_length.app_error", map[string]interface{}{"MaxLength": SITENAME_MAX_LENGTH}, "")
	}

}
