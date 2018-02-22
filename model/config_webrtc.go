package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

const (
	WEBRTC_CONFIG_FILE_PATH          = "./config/webrtc_config.json"
	WEBRTC_CONFIG_NAME               = "WEBRTC_SETTINGS"
	WEBRTC_SETTINGS_DEFAULT_STUN_URI = ""
	WEBRTC_SETTINGS_DEFAULT_TURN_URI = ""
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

func (s *WebrtcSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.GatewayWebsocketUrl == nil {
		s.GatewayWebsocketUrl = NewString("")
	}

	if s.GatewayAdminUrl == nil {
		s.GatewayAdminUrl = NewString("")
	}

	if s.GatewayAdminSecret == nil {
		s.GatewayAdminSecret = NewString("")
	}

	if s.StunURI == nil {
		s.StunURI = NewString(WEBRTC_SETTINGS_DEFAULT_STUN_URI)
	}

	if s.TurnURI == nil {
		s.TurnURI = NewString(WEBRTC_SETTINGS_DEFAULT_TURN_URI)
	}

	if s.TurnUsername == nil {
		s.TurnUsername = NewString("")
	}

	if s.TurnSharedKey == nil {
		s.TurnSharedKey = NewString("")
	}
}

func (ws *WebrtcSettings) isValid() *AppError {
	if *ws.Enable {
		if len(*ws.GatewayWebsocketUrl) == 0 || !IsValidWebsocketUrl(*ws.GatewayWebsocketUrl) {
			return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_ws_url.app_error", nil, "", http.StatusBadRequest)
		} else if len(*ws.GatewayAdminUrl) == 0 || !IsValidHttpUrl(*ws.GatewayAdminUrl) {
			return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_admin_url.app_error", nil, "", http.StatusBadRequest)
		} else if len(*ws.GatewayAdminSecret) == 0 {
			return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_gateway_admin_secret.app_error", nil, "", http.StatusBadRequest)
		} else if len(*ws.StunURI) != 0 && !IsValidTurnOrStunServer(*ws.StunURI) {
			return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_stun_uri.app_error", nil, "", http.StatusBadRequest)
		} else if len(*ws.TurnURI) != 0 {
			if !IsValidTurnOrStunServer(*ws.TurnURI) {
				return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_uri.app_error", nil, "", http.StatusBadRequest)
			}
			if len(*ws.TurnUsername) == 0 {
				return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_username.app_error", nil, "", http.StatusBadRequest)
			} else if len(*ws.TurnSharedKey) == 0 {
				return NewAppError("Config.IsValid", "model.config.is_valid.webrtc_turn_shared_key.app_error", nil, "", http.StatusBadRequest)
			}
		}
	}

	return nil
}

func webrtcConfigParser(f *os.File) (interface{}, error) {
	settings := &WebrtcSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("webrtc settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetWebRTCSettings() *WebrtcSettings {
	settings := utils.GetSettings(WEBRTC_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*WebrtcSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(WEBRTC_CONFIG_NAME, WEBRTC_CONFIG_FILE_PATH, true, webrtcConfigParser)
	if err != nil {
		return
	}
}
