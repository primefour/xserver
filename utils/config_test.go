package utils

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"io/ioutil"
	"sync"
	"testing"
)

type ConfigTest struct {
	config1 int    `json:"Config1"`
	config2 string `json:"Config2"`
	config3 bool   `json:"Config3"`
	config4 string `json:"Config4"`
}

var configTest = &ConfigTest{
	config1: 2,
	config2: "hello world",
	config3: true,
	config4: "true",
}

func configTestParser(buff []byte) {
	x := string(buff)
	l4g.Info(fmt.Sprintf("get a config buff is %s ", x))
	//err := json.Unmarshal(buff, configTest)
	//l4g.Info(fmt.Sprintf("get a config is %v %v ", configTest, err))

}

func TestConfig(t *testing.T) {
	sg := sync.WaitGroup{}
	sg.Add(1)
	_, err := NewXConfig("hello", "/home/crazyhorse/go/testGo/src/github.com/primefour/xserver", true, configTestParser)
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
	ferr := ioutil.WriteFile("./world.json", buff, 0664)
	if ferr != nil {
		t.Error("write file error %v ", err)
	}
	sg.Wait()
}
