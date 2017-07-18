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
	SaveConfig(fileName string, config interface{}) error //backup config
	LoadConfig(fileName string) (interface{}, error)      //load config from file
	VerifyConfig(config interface{}) bool                 //verify config file is correct
	SetDefault(config interface{}) interface{}            //set default after verify
	UpdateConfig(config interface{})                      //fnotify to update config
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
	mutex      sync.Mutex        //mutex for loading config file
	watcher    *fsnotify.Watcher //watcher for config file
	isWatch    bool
	FilePath   string //config file path
	AppName    string //app name and config name is appname.json
	ActivePath string //actually use by server
	Intf       ConfigIntf
}

func NewXConfig(app, filePath, isW) (*XConfig, error) {
	xc = new(XConfig{
		isWatch:  isW,
		AppName:  app,
		FilePath: filePath,
	})

	if !xc.checkExist() {
		return nil, l4g.Error(fmt.Sprintf("fail to load config for config file not exist"))
	}
	xc.initializeConfigWatch()
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

func (self *XConfig) initializeConfigWatch() {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if !self.isWatch {
		return
	}

	if self.watcher == nil {
		var err error
		self.watcher, err = fsnotify.NewWatcher()
		if err != nil {
			l4g.Error(fmt.Sprintf("Failed to watch config file at %v with err=%v", self.ActivePath, err.Error()))
		}

		go func() {
			configFile := self.ActivePath
			for {
				select {
				case event := <-self.watcher.Events:
					// we only care about the config file
					l4g.Info(fmt.Sprintf("notify file name %s ", event.Name))
					if filepath.Clean(event.Name) == configFile {
						if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
							l4g.Info(fmt.Sprintf("Config file watcher detected a change reloading %v", configFile))
							if configReadErr := viper.ReadInConfig(); configReadErr == nil {
								xj, err := self.Intf.LoadConfig(self)
								if err == nil {
									self.Intf.UpdateConfgi(xj)
								} else {
									l4g.Error(fmt.Sprintf("Failed to read while watching config file"))
								}
							} else {
								l4g.Error(fmt.Sprintf("Failed to read while watching config file at %v with err=%v", configFile, configReadErr.Error()))
							}
						}
					}
				case err := <-self.watcher.Errors:
					l4g.Error(fmt.Sprintf("Failed while watching config file at %v with err=%v", configFile, err.Error()))
				}
			}
		}()
	}
}
