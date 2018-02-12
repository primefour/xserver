package utils

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"os"
	"path/filepath"
	"sync"
)

type ConfigEntry struct {
	IsWatch  bool         //watch or not
	FilePath string       //config file dir
	Name     string       //config name
	Parser   ConfigParser //parse function
}

type ConfigParser func(f *os.File) (interface{}, error) //load config from file

//store config entries
var configEntries = map[string]*ConfigEntry{}

//cache settings
var configSettings = map[string]interface{}{}

var configMutex = sync.Mutex{}

func GetSettings(name string) interface{} {
	settings, ok := configSettings[name]
	if ok {
		return settings
	}
	return nil
}

func onFileUpdate(fullPath string) error {
	file, err := os.Open(fullPath)

	if err != nil {
		return l4g.Error(fmt.Sprintf("open config file fail %s ", fullPath))
	}

	defer file.Close()
	configEntry, ok := configEntries[fullPath]

	if !ok {
		return l4g.Warn(fmt.Sprintf("get config entry fail %s %v", fullPath, ok))
	}

	//buff, err := ioutil.ReadAll(file)
	if err == nil {
		setting, perr := configEntry.Parser(file)
		if perr == nil {
			configSettings[configEntry.Name] = setting
			return nil
		} else {
			l4g.Error("parser config file %s failed %v ", fullPath, perr)
		}
	}
	return l4g.Error("%s parser failed", fullPath)
}

func AddConfigEntry(name, filePath string, isW bool, parser ConfigParser) (*ConfigEntry, error) {
	configMutex.Lock()
	defer configMutex.Unlock()

	entry := &ConfigEntry{
		IsWatch:  isW,
		Name:     name,
		FilePath: filePath,
		Parser:   parser,
	}

	if !entry.checkExist() {
		return nil, l4g.Error(fmt.Sprintf("fail to load config for config file not exist"))
	}
	configEntries[entry.FilePath] = entry

	if entry.IsWatch {
		AddFileWatch(entry.FilePath, onFileUpdate)
	}

	err := onFileUpdate(entry.FilePath)

	return entry, err
}

func (self *ConfigEntry) checkExist() bool {
	//check
	if len(self.Name) == 0 {
		l4g.Error("Setting name is empty")
		return false
	}

	//check
	if len(self.FilePath) != 0 {
		if _, err := os.Stat(self.FilePath); err == nil {
			self.FilePath, _ = filepath.Abs(self.FilePath)
			self.FilePath = filepath.Clean(self.FilePath)
			l4g.Info(fmt.Sprintf("%s use config file is %s ", self.Name, self.FilePath))
			return true
		}
	}

	l4g.Error(fmt.Sprintf("%s settings no config file path", self.Name))
	return false
}
