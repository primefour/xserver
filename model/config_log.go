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

func (self *LogSettings) isValid() *utils.AppError {
	if self.FileLevel != "DEBUG" || self.FileLevel != "INFO" ||
		self.FileLevel != "WARN" || self.FileLevel != "ERROR" {
		return utils.NewLocAppError("Config.IsValid", "utils.logconfig.is_valid.level_error", nil, "")
	}
	return nil
}

func (self *LogSettings) setDefaults() {
	if self.EnableDiagnostics == nil {
		self.EnableDiagnostics = new(bool)
		*self.EnableDiagnostics = true
	}
	self.EnableConsole = true
	self.ConsoleLevel = "DEBUG"
	self.EnableFile = true
	self.FileLevel = "DEBUG"
}

func logConfigParser(f *os.File) (interface{}, error) {
	settings := &LogSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.setDefaults()
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
