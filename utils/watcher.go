package utils

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"sync"
)

type ConfigUpdateFpn func(file string)

var fileWatcher *fsnotify.Watcher
var watcher_mutex sync.Mutex = sync.Mutex{}
var fileNameMap map[string]ConfigUpdateFpn = map[string]ConfigUpdateFpn{}
var dirMap map[string]int = map[string]int{}
var onceNotifyInit = sync.Once{}

func EnableConfigFromEnviromentVars() {
	viper.SetEnvPrefix("xs")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func init() {
	fw, err := fsnotify.NewWatcher()
	fileWatcher = nil
	if err == nil {
		fileWatcher = fw
	} else {
		l4g.Error("create config file watcher fail")
		return
	}
}

func watcherOnce() {
	go WatcherNotify()
}

func WatcherNotify() {
	if fileWatcher == nil {
		l4g.Warn("file watcher object is nil ")
		return
	}

	l4g.Info("enter go for loop and monitor")
	for {
		select {
		case event := <-fileWatcher.Events:
			// we only care about the config file
			cfn := filepath.Clean(event.Name)
			fpn, ok := fileNameMap[cfn]
			if ok && fpn != nil {
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					l4g.Info(fmt.Sprintf("Config file watcher detected a change reloading %v", cfn))
					fpn(cfn)
				}
			}
		case err := <-fileWatcher.Errors:
			l4g.Error(fmt.Sprintf("Failed while watching config file with err=%v", err.Error()))
		}
	}
}

func AddConfigWatch(file string, updateFpn ConfigUpdateFpn) {
	watcher_mutex.Lock()
	defer watcher_mutex.Unlock()
	if fileWatcher != nil {
		configDir, _ := filepath.Split(file)
		_, ok := dirMap[configDir]
		if ok {
			dirMap[configDir]++
		} else {
			dirMap[configDir] = 1
			fileWatcher.Add(configDir)
		}
	}
	fileNameMap[file] = updateFpn
	onceNotifyInit.Do(watcherOnce)
}

func RemoveConfigWatch(file string) {
	watcher_mutex.Lock()
	defer watcher_mutex.Unlock()
	if fileWatcher != nil {
		configDir, _ := filepath.Split(file)
		value, ok := dirMap[configDir]
		if ok {
			if value == 1 {
				delete(dirMap, configDir)
				fileWatcher.Remove(configDir)
			} else {
				dirMap[configDir]--
			}
		} else {
			return
		}
	}
	delete(fileNameMap, file)
}
