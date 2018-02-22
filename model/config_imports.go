package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	SERVICE_GITLAB           = "gitlab"
	SERVICE_GOOGLE           = "google"
	SERVICE_OFFICE365        = "office365"
	IMPORTS_CONFIG_FILE_PATH = "./config/imports_config.json"
	IMPORTS_CONFIG_NAME      = "IMPORTS_SETTINGS"
)

type SSOSettings struct {
	Enable          bool
	Secret          string
	Id              string
	Scope           string
	AuthEndpoint    string
	TokenEndpoint   string
	UserApiEndpoint string
}

type ImportsSettings struct {
	GitLabSettings    SSOSettings
	GoogleSettings    SSOSettings
	Office365Settings SSOSettings
}

func (o *ImportsSettings) GetSSOService(service string) *SSOSettings {
	switch service {
	case SERVICE_GITLAB:
		return &o.GitLabSettings
	case SERVICE_GOOGLE:
		return &o.GoogleSettings
	case SERVICE_OFFICE365:
		return &o.Office365Settings
	}
	return nil
}

func (o *ImportsSettings) Sanitize() {
	if len(o.GitLabSettings.Secret) > 0 {
		o.GitLabSettings.Secret = FAKE_SETTING
	}
}

func importsConfigParser(f *os.File) (interface{}, error) {
	settings := &ImportsSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	l4g.Debug("imports settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetImportsSettings() *ImportsSettings {
	settings := utils.GetSettings(IMPORTS_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*ImportsSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(IMPORTS_CONFIG_NAME, IMPORTS_CONFIG_FILE_PATH, true, importsConfigParser)
	if err != nil {
		return
	}
}
