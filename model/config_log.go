package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	LOG_CONFIG_FILE_PATH = "./config/log_config.json"
	LOG_CONFIG_NAME      = "LOG_SETTINGS"
)

type LogSettings struct {
	EnableConsole          bool
	ConsoleLevel           string
	EnableFile             bool
	FileLevel              string
	FileFormat             string
	FileLocation           string
	EnableWebhookDebugging bool
	EnableDiagnostics      *bool
}

func (s *LogSettings) SetDefaults() {
	if s.EnableDiagnostics == nil {
		s.EnableDiagnostics = NewBool(true)
	}
}

func logConfigParser(f *os.File) (interface{}, error) {
	settings := &LogSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("log settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetLogSettings() *LogSettings {
	settings := utils.GetSettings(LOG_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*LogSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(LOG_CONFIG_NAME, LOG_CONFIG_FILE_PATH, true, logConfigParser)
	if err != nil {
		return
	}
}
