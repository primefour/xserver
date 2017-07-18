package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type ConfigIntf interface {
	LoadConfig(buff []byte) (interface{}, error) //load config from file
	VerifyConfig(config interface{}) bool        //verify config file is correct
	SetDefault(config interface{}) interface{}   //set default after verify
}

func LoadConfig(config *XConfig) (interface{}, error) {
	if config == nil {
		return l4g.Error(fmt.Sprintf("fail to load config for config is nil "))
	}
	cf, err := config.Intf.LoadConfig(config.ActivePath)
	if err != nil {
		return l4g.Error(fmt.Sprintf("fail to load config for config file format"))
	}
	if config.Intf.VerifyConfig(cf) {
		return l4g.Error(fmt.Sprintf("fail to load config for config verify"))
	}
	return config.Intf.SetDefault(config)
}

//default is ./config/
func findConfigFile(fileName string) string {
	if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	}

	return fileName
}

type XConfig struct {
	IsWatch    bool
	FileDir    string //config file dir
	AppName    string //app name and config name is appname.json
	ActivePath string //actually use by server
	Intf       ConfigIntf
}

func NewXConfig(app, dir, isW, intf ConfigIntf) (*XConfig, error) {
	xc = new(XConfig{
		IsWatch:  isW,
		AppName:  app,
		FilePath: dir,
		Intf:     intf,
	})
	if !xc.checkExist() {
		return nil, l4g.Error(fmt.Sprintf("fail to load config for config file not exist"))
	}
	return xc, nil
}

func (self *XConfig) checkExist() bool {
	//check
	if len(self.AppName) == 0 {
		return false
	}
	//check
	var configPath string
	if len(self.FilePath) != 0 {
		configPath = self.FilePath + "/" + self.AppName + ".json"
		if _, err := os.Stat(configPath); err == nil {
			configPath, _ = filepath.Abs(configPath)
			configPath := filepath.Clean(configPath)
			l4g.Info(fmt.Sprintf("%s use config file is %s ", self.AppName, configPath))
			self.ActivePath = configPath
			return true
		}
	} else {
		configPath = findConfigFile(self.AppName + ".json")
		if _, err := os.Stat(configPaht); err == nil {
			configPath := filepath.Clean(configPath)
			l4g.Info(fmt.Sprintf("%s use config file is %s ", self.AppName, configPath))
			self.ActivePath = configPath
			return true
		}
	}
	return false
}
