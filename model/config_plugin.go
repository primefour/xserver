package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	PLUGIN_CONFIG_FILE_PATH                  = "./config/plugin_config.json"
	PLUGIN_CONFIG_NAME                       = "PLUGIN_SETTINGS"
	PLUGIN_SETTINGS_DEFAULT_DIRECTORY        = "./plugins"
	PLUGIN_SETTINGS_DEFAULT_CLIENT_DIRECTORY = "./client/plugins"
)

type PluginState struct {
	Enable bool
}

type PluginSettings struct {
	Enable          *bool
	EnableUploads   *bool
	Directory       *string
	ClientDirectory *string
	Plugins         map[string]interface{}
	PluginStates    map[string]*PluginState
}

func (s *PluginSettings) SetDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(true)
	}

	if s.EnableUploads == nil {
		s.EnableUploads = NewBool(false)
	}

	if s.Directory == nil {
		s.Directory = NewString(PLUGIN_SETTINGS_DEFAULT_DIRECTORY)
	}

	if *s.Directory == "" {
		*s.Directory = PLUGIN_SETTINGS_DEFAULT_DIRECTORY
	}

	if s.ClientDirectory == nil {
		s.ClientDirectory = NewString(PLUGIN_SETTINGS_DEFAULT_CLIENT_DIRECTORY)
	}

	if *s.ClientDirectory == "" {
		*s.ClientDirectory = PLUGIN_SETTINGS_DEFAULT_CLIENT_DIRECTORY
	}

	if s.Plugins == nil {
		s.Plugins = make(map[string]interface{})
	}

	if s.PluginStates == nil {
		s.PluginStates = make(map[string]*PluginState)
	}
}

func pluginConfigParser(f *os.File) (interface{}, error) {
	settings := &PluginSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("plugin settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetPluginSettings() *PluginSettings {
	settings := utils.GetSettings(PLUGIN_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*PluginSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(PLUGIN_CONFIG_NAME, PLUGIN_CONFIG_FILE_PATH, true, pluginConfigParser)
	if err != nil {
		return
	}
}
