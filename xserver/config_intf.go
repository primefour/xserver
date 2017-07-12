package xserver

import (
	l4g "github.com/alecthomas/log4go"
)

type ConfigIntf interface {
	SaveConfig(path string, config interface{}) error
	LoadConfig(path string) (error, interface{})
	IsValid(config_value interface{}) bool
	SetDefault(config_value interface{}) interface{}
}

func GetConfig(path string, config ConfigIntf) (configv interface{}, err error) {
	if len(path) == 0 || config == nil {
		err = l4g.Error("path for config is null")
		return nil, err
	}

	if err, config_value := config.LoadConfig(path); err != nil {
		l4g.Error("load config fail path is " + path)
		return nil, err
	} else {
		config_value = config.SetDefault(config_value)
		if !config.IsValid(config_value) {
			err = l4g.Error("config value is invalidate ")
			return nil, err
		}
		return config_value, nil
	}
}
