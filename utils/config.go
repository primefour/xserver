package utils

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type ConfigParserFpn func(buff []byte) //load config from file
var fileXConfigMap = map[string]*XConfig{}
var config_mutex = sync.Mutex{}

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
	Parser     ConfigParserFpn
}

func onFileUpdate(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		l4g.Error(fmt.Sprintf("open config file fail %s ", fileName))
		return
	}

	defer file.Close()
	xc, ok := fileXConfigMap[fileName]

	if !ok {
		l4g.Warn(fmt.Sprintf("get xconfig of config file fail %s %v", fileName, ok))
		return
	}

	buff, err := ioutil.ReadAll(file)
	xc.Parser(buff)
}

func NewXConfig(app, dir string, isW bool, parser ConfigParserFpn) (*XConfig, error) {
	config_mutex.Lock()
	defer config_mutex.Unlock()

	xc := &XConfig{
		IsWatch: isW,
		AppName: app,
		FileDir: dir,
		Parser:  parser,
	}
	if !xc.checkExist() {
		return nil, l4g.Error(fmt.Sprintf("fail to load config for config file not exist"))
	}

	if xc.IsWatch {
		AddFileWatch(xc.ActivePath, onFileUpdate)
	}
	l4g.Info("xc.ActivePath is %s ", xc.ActivePath)
	fileXConfigMap[xc.ActivePath] = xc
	return xc, nil
}

func (self *XConfig) checkExist() bool {
	//check
	if len(self.AppName) == 0 {
		return false
	}
	//check
	var configPath string
	if len(self.FileDir) != 0 {
		configPath = self.FileDir + "/" + self.AppName + ".json"
		if _, err := os.Stat(configPath); err == nil {
			configPath, _ = filepath.Abs(configPath)
			configPath := filepath.Clean(configPath)
			l4g.Info(fmt.Sprintf("%s use config file is %s ", self.AppName, configPath))
			self.ActivePath = configPath
			return true
		}
	} else {
		configPath = findConfigFile(self.AppName + ".json")
		if _, err := os.Stat(configPath); err == nil {
			configPath := filepath.Clean(configPath)
			l4g.Info(fmt.Sprintf("%s use config file is %s ", self.AppName, configPath))
			self.ActivePath = configPath
			return true
		}
	}
	return false
}

func (self *XConfig) UpdateForce() {
	onFileUpdate(self.ActivePath)
}
