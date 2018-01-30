package utils

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
	"sync"
)

type onFileUpdateFunc func(fullPath string) error

var fileWatcher *fsnotify.Watcher = nil
var watchMutex sync.Mutex = sync.Mutex{}
var onFilesUpdateMap map[string]onFileUpdateFunc = map[string]onFileUpdateFunc{}
var onceNotifyInit = sync.Once{}

func init() {
	fw, err := fsnotify.NewWatcher()
	if err == nil {
		fileWatcher = fw
	} else {
		l4g.Error("create file watcher fail")
		return
	}
}

func watcherOnce() {
	go watchNotify()
}

func watchNotify() {
	if fileWatcher == nil {
		l4g.Warn("file watcher object is nil ")
		return
	}

	l4g.Info("enter go for loop and monitor")
	for {
		select {
		case event := <-fileWatcher.Events:
			watchMutex.Lock()
			// we only care about the register file
			filePath := filepath.Clean(event.Name)
			fileUpdate, ok := onFilesUpdateMap[filePath]
			if ok && fileUpdate != nil {
				if event.Op&fsnotify.Write == fsnotify.Write ||
					event.Op&fsnotify.Create == fsnotify.Create {
					l4g.Info(fmt.Sprintf("file watcher detected a change reloading %v", filePath))
					fileUpdate(filePath)
				}
			}
			watchMutex.Unlock()
		case err := <-fileWatcher.Errors:
			l4g.Error(fmt.Sprintf("Failed while watching file with err=%v", err.Error()))
		}
	}
}

func AddFileWatch(file string, updateFunc onFileUpdateFunc) {
	watchMutex.Lock()
	defer watchMutex.Unlock()
	file, _ = filepath.Abs(file)

	if fileWatcher != nil {
		fileWatcher.Add(file)
	}
	l4g.Debug("add watcher for file %s ", file)
	onFilesUpdateMap[file] = updateFunc
	onceNotifyInit.Do(watcherOnce)
}

func RemoveFileWatch(file string) {
	watchMutex.Lock()
	defer watchMutex.Unlock()
	file, _ = filepath.Abs(file)
	if fileWatcher != nil {
		fileWatcher.Remove(file)
	}
	delete(onFilesUpdateMap, file)
}
