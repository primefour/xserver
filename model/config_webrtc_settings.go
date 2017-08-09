package model

import (
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
	if self.Enable == nil {
		self.Enable = new(bool)
		*self.Enable = false
	}

	if self.GatewayWebsocketUrl == nil {
		self.GatewayWebsocketUrl = new(string)
		*self.GatewayWebsocketUrl = ""
	}

	if self.GatewayAdminUrl == nil {
		self.GatewayAdminUrl = new(string)
		*self.GatewayAdminUrl = ""
	}

	if self.GatewayAdminSecret == nil {
		self.GatewayAdminSecret = new(string)
		*self.GatewayAdminSecret = ""
	}

	if self.StunURI == nil {
		self.StunURI = new(string)
		*self.StunURI = WEBRTC_SETTINGS_DEFAULT_STUN_URI
	}

	if self.TurnURI == nil {
		self.TurnURI = new(string)
		*self.TurnURI = WEBRTC_SETTINGS_DEFAULT_TURN_URI
	}

	if self.TurnUsername == nil {
		self.TurnUsername = new(string)
		*self.TurnUsername = ""
	}

	if self.TurnSharedKey == nil {
		self.TurnSharedKey = new(string)
		*self.TurnSharedKey = ""
	}
}

func (self *WebrtcSettings) IsValidate() *utils.AppError {
	if *self.Enable {
		if len(*self.GatewayWebsocketUrl) == 0 || !utils.IsValidWebsocketUrl(*self.GatewayWebsocketUrl) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_ws_url.app_error", nil, "")
		} else if len(*self.GatewayAdminUrl) == 0 || !utils.IsValidHttpUrl(*self.GatewayAdminUrl) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_admin_url.app_error", nil, "")
		} else if len(*self.GatewayAdminSecret) == 0 {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_admin_secret.app_error", nil, "")
		} else if len(*self.StunURI) != 0 && !utils.IsValidTurnOrStunServer(*self.StunURI) {
			return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_stun_uri.app_error", nil, "")
		} else if len(*self.TurnURI) != 0 {
			if !utils.IsValidTurnOrStunServer(*self.TurnURI) {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_uri.app_error", nil, "")
			}
			if len(*self.TurnUsername) == 0 {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_username.app_error", nil, "")
			} else if len(*self.TurnSharedKey) == 0 {
				return utils.NewLocAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_shared_key.app_error", nil, "")
			}

		}
	}

	return nil
}
