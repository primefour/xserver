package utils

import (
	"fmt"
	"sync"
	"testing"
)

func update(filePath string) {
	fmt.Printf("update get a notify %s \n", filePath)
}

func TestWatch(T *testing.T) {
	dir := "/home/crazyhorse/go/testGo/src/github.com/primefour/xserver/"
	fileName := "hello.w"
	filePath := dir + fileName
	var wg = sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("monitor file %s \n", filePath)
	AddFileWatch(filePath, update)
	go func() {
		WatcherNotify()
		wg.Done()
	}()
	RemoveFileWatch(filePath)

	AddFileWatch(filePath, update)
	wg.Wait()
}
