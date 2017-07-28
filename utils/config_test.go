package utils

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"
)

type ConfigTest struct {
	Configx int    `json:"Config1"`
	Configy string `json:"Config2"`
	Configz bool   `json:"Config3"`
	Configi string `json:"Config4"`
}

var configTest = &ConfigTest{
	Configx: 2,
	Configy: "hello world",
	Configz: true,
	Configi: "true",
}

var configTest2 = &ConfigTest{}

func configTestParser(buff []byte) {
	x := string(buff)
	l4g.Info(fmt.Sprintf("get a config buff is %s ", x))
	err := json.Unmarshal(buff, configTest2)
	l4g.Info(fmt.Sprintf("get a config is %v %v ", configTest2, err))
}

var testpwd = "./"

func say(s string) {
	fmt.Println(s)
	time.Sleep(time.Second * 500)
}

func TestConfig(t *testing.T) {
	sg := sync.WaitGroup{}
	sg.Add(1)
	pwd, _ := os.Getwd()

	fmt.Printf("pwd = %s pid %d \n", pwd, os.Getpid())

	_, err := NewXConfig("hello", testpwd, true, configTestParser)
	if err != nil {
		t.Error(" fail for new Xconfig ")
	}
	l4g.Info("config test is %v ", configTest)
	buff, err := json.Marshal(&configTest)
	if err != nil {
		t.Error("test json marshal fail")
	}
	strc := string(buff)
	l4g.Info("buff info %s ", strc)
	ferr := ioutil.WriteFile(testpwd+"world.json", buff, 0664)
	if ferr != nil {
		t.Error("write file error %v ", err)
	}
	sg.Wait()
}
