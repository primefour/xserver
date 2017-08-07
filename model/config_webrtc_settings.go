package model

import (
	"encoding/json"
	"github.com/primefour/xserver/utils"
)

type WebrtcSettings struct {
	Enable              *bool
	GatewayWebsocketUrl *string
	GatewayAdminUrl     *string
	GatewayAdminSecret  *string
	StunURI             *string
	TurnURI             *string
	TurnUsername        *string
	TurnSharedKey       *string
}

func (self *WebrtcSettings) setDefault() {
	if self.WebrtcSettings.Enable == nil {
		self.WebrtcSettings.Enable = new(bool)
		*self.WebrtcSettings.Enable = false
	}

	if self.WebrtcSettings.GatewayWebsocketUrl == nil {
		self.WebrtcSettings.GatewayWebsocketUrl = new(string)
		*self.WebrtcSettings.GatewayWebsocketUrl = ""
	}

	if self.WebrtcSettings.GatewayAdminUrl == nil {
		self.WebrtcSettings.GatewayAdminUrl = new(string)
		*self.WebrtcSettings.GatewayAdminUrl = ""
	}

	if self.WebrtcSettings.GatewayAdminSecret == nil {
		self.WebrtcSettings.GatewayAdminSecret = new(string)
		*self.WebrtcSettings.GatewayAdminSecret = ""
	}

	if self.WebrtcSettings.StunURI == nil {
		self.WebrtcSettings.StunURI = new(string)
		*self.WebrtcSettings.StunURI = WEBRTC_SETTINGS_DEFAULT_STUN_URI
	}

	if self.WebrtcSettings.TurnURI == nil {
		self.WebrtcSettings.TurnURI = new(string)
		*self.WebrtcSettings.TurnURI = WEBRTC_SETTINGS_DEFAULT_TURN_URI
	}

	if self.WebrtcSettings.TurnUsername == nil {
		self.WebrtcSettings.TurnUsername = new(string)
		*self.WebrtcSettings.TurnUsername = ""
	}

	if self.WebrtcSettings.TurnSharedKey == nil {
		self.WebrtcSettings.TurnSharedKey = new(string)
		*self.WebrtcSettings.TurnSharedKey = ""
	}
}

func (self *WebrtcSettings) IsValidate() *utils.AppError {
	if *self.WebrtcSettings.Enable {
		if len(*self.WebrtcSettings.GatewayWebsocketUrl) == 0 || !utils.IsValidWebsocketUrl(*self.WebrtcSettings.GatewayWebsocketUrl) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_ws_url.app_error", nil, "")
		} else if len(*self.WebrtcSettings.GatewayAdminUrl) == 0 || !utils.IsValidHttpUrl(*self.WebrtcSettings.GatewayAdminUrl) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_admin_url.app_error", nil, "")
		} else if len(*self.WebrtcSettings.GatewayAdminSecret) == 0 {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_admin_secret.app_error", nil, "")
		} else if len(*self.WebrtcSettings.StunURI) != 0 && !utils.IsValidTurnOrStunServer(*self.WebrtcSettings.StunURI) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_stun_uri.app_error", nil, "")
		} else if len(*self.WebrtcSettings.TurnURI) != 0 {
			if !utils.IsValidTurnOrStunServer(*self.WebrtcSettings.TurnURI) {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_uri.app_error", nil, "")
			}
			if len(*self.WebrtcSettings.TurnUsername) == 0 {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_username.app_error", nil, "")
			} else if len(*self.WebrtcSettings.TurnSharedKey) == 0 {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_shared_key.app_error", nil, "")
			}

		}
	}

	return nil
}
