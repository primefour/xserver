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

type FileUpdateFpn func(file string)

type DirUpdateFpn func(file string)

var fileWatcher *fsnotify.Watcher = nil
var watcher_mutex sync.Mutex = sync.Mutex{}
var fileNameMap map[string]FileUpdateFpn = map[string]FileUpdateFpn{}
var dirNameMap map[string]DirUpateFpn = map[string]DirUpdateFpn{}
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
		l4g.Error("create file watcher fail")
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
			// we only care about the register file
			cfn := filepath.Clean(event.Name)
			fpn, ok := fileNameMap[cfn]
			if ok && fpn != nil {
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					l4g.Info(fmt.Sprintf("file watcher detected a change reloading %v", cfn))
					fpn(cfn)
				}
			}
			//we only care about the directory we register
			edir, _ := filepath.Split(event.Name)
			dpn, eok := dirNameMap[edir]
			if eok && dpn != nil {
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					l4g.Info(fmt.Sprintf("dir watcher detected a change reloading %v", edir))
					dpn(edir)
				}
			}
		case err := <-fileWatcher.Errors:
			l4g.Error(fmt.Sprintf("Failed while watching file with err=%v", err.Error()))
		}
	}
}

func AddFileWatch(file string, updateFpn FileUpdateFpn) {
	watcher_mutex.Lock()
	defer watcher_mutex.Unlock()
	file, _ = filepath.Abs(file)
	if fileWatcher != nil {
		fileDir, _ := filepath.Split(file)
		_, ok := dirMap[fileDir]
		if ok {
			dirMap[fileDir]++
		} else {
			dirMap[fileDir] = 1
			fileWatcher.Add(fileDir)
		}
	}
	fileNameMap[file] = updateFpn
	onceNotifyInit.Do(watcherOnce)
}

func RemoveFileWatch(file string) {
	watcher_mutex.Lock()
	defer watcher_mutex.Unlock()
	file, _ = filepath.Abs(file)
	if fileWatcher != nil {
		fileDir, _ := filepath.Split(file)
		value, ok := dirMap[fileDir]
		if ok {
			if value == 1 {
				delete(dirMap, fileDir)
				fileWatcher.Remove(fileDir)
			} else {
				dirMap[fileDir]--
			}
		} else {
			return
		}
	}
	delete(fileNameMap, file)
}

func AddDirWatch(dir string, updatefpn DirUpdateFpn) {
	watcher_mutex.Lock()
	defer watcher_mutex.Unlock()
	dir, _ = filepath.Abs(dir)
	var fileDir string
	if fileWatcher != nil {
		fileDir, _ = filepath.Split(dir)
		_, ok := dirMap[fileDir]
		if ok {
			dirMap[fileDir]++
		} else {
			dirMap[fileDir] = 1
			fileWatcher.Add(fileDir)
		}
	}
	dirNameMap[fileDir] = updateFpn
	onceNotifyInit.Do(watcherOnce)
}

func RemoveDirWatch(dir string) {
	watcher_mutex.Lock()
	defer watcher_mutex.Unlock()
	dir, _ = filepath.Abs(dir)
	var fileDir string
	if fileWatcher != nil {
		fileDir, _ = filepath.Split(dir)
		value, ok := dirMap[fileDir]
		if ok {
			if value == 1 {
				delete(dirMap, fileDir)
				fileWatcher.Remove(fileDir)
			} else {
				dirMap[fileDir]--
			}
		} else {
			return
		}
	}
	delete(dirNameMap, fileDir)
}
