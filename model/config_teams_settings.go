package model

import (
	"github.com/primefour/xserver/utils"
)

/*******************************************************************************
 **      Filename: model/config_teams_settings.go
 **        Author: crazyhorse
 **   Description: ---
 **        Create: 2017-08-09 10:37:25
 ** Last Modified: 2017-08-09 10:37:25
 ******************************************************************************/

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
	if self.EnableCustomBrand == nil {
		self.EnableCustomBrand = new(bool)
		*self.EnableCustomBrand = false
	}

	if self.CustomBrandText == nil {
		self.CustomBrandText = new(string)
		*self.CustomBrandText = TEAM_SETTINGS_DEFAULT_CUSTOM_BRAND_TEXT
	}

	if self.CustomDescriptionText == nil {
		self.CustomDescriptionText = new(string)
		*self.CustomDescriptionText = TEAM_SETTINGS_DEFAULT_CUSTOM_DESCRIPTION_TEXT
	}

	if self.EnableOpenServer == nil {
		self.EnableOpenServer = new(bool)
		*self.EnableOpenServer = false
	}

	if self.RestrictDirectMessage == nil {
		self.RestrictDirectMessage = new(string)
		*self.RestrictDirectMessage = DIRECT_MESSAGE_ANY
	}

	if self.RestrictTeamInvite == nil {
		self.RestrictTeamInvite = new(string)
		*self.RestrictTeamInvite = PERMISSIONS_ALL
	}

	if self.RestrictPublicChannelManagement == nil {
		self.RestrictPublicChannelManagement = new(string)
		*self.RestrictPublicChannelManagement = PERMISSIONS_ALL
	}

	if self.RestrictPrivateChannelManagement == nil {
		self.RestrictPrivateChannelManagement = new(string)
		*self.RestrictPrivateChannelManagement = PERMISSIONS_ALL
	}

	if self.RestrictPublicChannelCreation == nil {
		self.RestrictPublicChannelCreation = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		if *self.RestrictPublicChannelManagement == PERMISSIONS_CHANNEL_ADMIN {
			*self.RestrictPublicChannelCreation = PERMISSIONS_TEAM_ADMIN
		} else {
			*self.RestrictPublicChannelCreation = *self.RestrictPublicChannelManagement
		}
	}

	if self.RestrictPrivateChannelCreation == nil {
		self.RestrictPrivateChannelCreation = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		if *self.RestrictPrivateChannelManagement == PERMISSIONS_CHANNEL_ADMIN {
			*self.RestrictPrivateChannelCreation = PERMISSIONS_TEAM_ADMIN
		} else {
			*self.RestrictPrivateChannelCreation = *self.RestrictPrivateChannelManagement
		}
	}

	if self.RestrictPublicChannelDeletion == nil {
		self.RestrictPublicChannelDeletion = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		*self.RestrictPublicChannelDeletion = *self.RestrictPublicChannelManagement
	}

	if self.RestrictPrivateChannelDeletion == nil {
		self.RestrictPrivateChannelDeletion = new(string)
		// If this setting does not exist, assume migration from <3.6, so use management setting as default.
		*self.RestrictPrivateChannelDeletion = *self.RestrictPrivateChannelManagement
	}

	if self.RestrictPrivateChannelManageMembers == nil {
		self.RestrictPrivateChannelManageMembers = new(string)
		*self.RestrictPrivateChannelManageMembers = PERMISSIONS_ALL
	}

	if self.UserStatusAwayTimeout == nil {
		self.UserStatusAwayTimeout = new(int64)
		*self.UserStatusAwayTimeout = TEAM_SETTINGS_DEFAULT_USER_STATUS_AWAY_TIMEOUT
	}

	if self.MaxChannelsPerTeam == nil {
		self.MaxChannelsPerTeam = new(int64)
		*self.MaxChannelsPerTeam = 2000
	}

	if self.MaxNotificationsPerChannel == nil {
		self.MaxNotificationsPerChannel = new(int64)
		*self.MaxNotificationsPerChannel = 1000
	}
}

func (self *TeamSettings) IsValidate() *utils.AppError {

	if self.MaxUsersPerTeam <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_users.app_error", nil, "")
	}

	if *self.MaxChannelsPerTeam <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_channels.app_error", nil, "")
	}

	if *self.MaxNotificationsPerChannel <= 0 {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.max_notify_per_channel.app_error", nil, "")
	}

	if !(*self.RestrictDirectMessage == DIRECT_MESSAGE_ANY || *self.RestrictDirectMessage == DIRECT_MESSAGE_TEAM) {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.restrict_direct_message.app_error", nil, "")
	}

	if len(self.SiteName) > SITENAME_MAX_LENGTH {
		return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.sitename_length.app_error", map[string]interface{}{"MaxLength": SITENAME_MAX_LENGTH}, "")
	}
	return nil
}
