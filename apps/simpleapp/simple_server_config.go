package simpleapp

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/utils"
)

var configFileName = "simple_server"
var configFilePath = "/home/crazyhorse/go/testGo/src/github.com/primefour/xserver/apps/simpleapp"
var SServConfig SimpleServConfig = SimpleServConfig{}
var ssxconfig *utils.XConfig = nil

type SimpleServConfig struct {
	ServiceSettings      model.ServiceSettings
	SqlSettings          model.SqlSettings
	FileSettings         model.FileSettings
	RateLimitSettings    model.RateLimitSettings
	LocalizationSettings model.LocalizationSettings
	EmailSettings        model.EmailSettings
}

func (self *SimpleServConfig) SetDefault() {
	self.ServiceSettings.SetDefault()
	self.SqlSettings.SetDefault()
	self.FileSettings.SetDefault()
	self.RateLimitSettings.SetDefault()
	self.LocalizationSettings.SetDefault()
	self.EmailSettings.SetDefault()
}

func (self *SimpleServConfig) ToJson() string {
	b, err := json.Marshal(self)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

/*
func (self *SimpleServConfig) GetSSOService(service string) *SSOSettings {
	return nil
}
*/

func SimpleServConfigFromJson(buff []byte) {
	x := string(buff)
	l4g.Info(fmt.Sprintf("get a config buff is %s ", x))
	SServConfig.SetDefault()
	err := json.Unmarshal(buff, &SServConfig)
	if err != nil {
		l4g.Error(fmt.Sprintf("parse simple server config file fail %v ", err))
		panic(err)
	}
	l4g.Info(fmt.Sprintf("get a config is %v %v ", SServConfig, err))
}

func init() {
	ssxconfig, err := utils.NewXConfig(configFileName, configFilePath, true, SimpleServConfigFromJson)
	if err != nil {
		l4g.Error("create simple server xconfig fail")
		panic(err)
	}
	ssxconfig.UpdateForce()
}
