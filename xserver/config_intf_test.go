package xserver

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"os"
	"testing"
)

type XLogSetting struct {
	EnableConsole          bool
	ConsoleLevel           string
	EnableFile             bool
	FileLevel              string
	FileFormat             string
	FileLocation           string
	EnableWebhookDebugging bool
	EnableDiagnostics      *bool
}

var LogConfig = XLogSetting{}

func (self *XLogSetting) SaveConfig(path string, config interface{}) error {
	return nil
}

func (self *XLogSetting) LoadConfig(path string) (error, interface{}) {
	file, err := os.Open(path)
	if err != nil {
		xerr := l4g.Error("open file path fail " + path)
		return xerr, nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(self)
	if err == nil {
		return nil, *self
	} else {
		return err, nil
	}

}

func (self *XLogSetting) IsValid(config_value interface{}) bool {
	return true
}

func (self *XLogSetting) SetDefault(config_value interface{}) interface{} {
	xj, ok := config_value.(XLogSetting)
	if ok {
		xj.EnableConsole = false
	} else {
		fmt.Println("is not *xlogsetting type")
	}
	return xj
}

func changeValue(x interface{}) {
	xv, ok := x.(XLogSetting)
	if ok {
		fmt.Println("is xlogsetting type")
		xv.EnableConsole = false
	} else {
		fmt.Println("is not xlogsetting type ")
	}
}

func TestGetConfig(t *testing.T) {
	config, err := GetConfig("./log_config.json", &LogConfig)
	if err != nil {
		t.Error(" parse config file fail ")
	} else {
		fmt.Printf("%v ", config)
	}
}
