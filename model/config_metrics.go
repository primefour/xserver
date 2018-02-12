package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	//metrics config
	METRICS_CONFIG_FILE_PATH = "./config/metrics_config.json"
	METRICS_CONFIG_NAME      = "METRICS_SETTINGS"
)

type MetricsSettings struct {
	Enable           *bool
	BlockProfileRate *int
	ListenAddress    *string
}

func (s *MetricsSettings) SetDefaults() {
	if s.ListenAddress == nil {
		s.ListenAddress = NewString(":8067")
	}

	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.BlockProfileRate == nil {
		s.BlockProfileRate = NewInt(0)
	}
}

func metricsConfigParser(f *os.File) (interface{}, error) {
	settings := &MetricsSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("metrics settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetMetricsSettings() *MetricsSettings {
	settings := utils.GetSettings(METRICS_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*MetricsSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(METRICS_CONFIG_NAME, METRICS_CONFIG_FILE_PATH, true, metricsConfigParser)
	if err != nil {
		return
	}
}
