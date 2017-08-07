package model

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
)

type SimpleServerConfig struct {
	ServiceSettings   ServiceSettings
	TeamSettings      TeamSettings
	SqlSettings       SqlSettings
	LogSettings       LogSettings
	PasswordSettings  PasswordSettings
	FileSettings      FileSettings
	EmailSettings     EmailSettings
	RateLimitSettings RateLimitSettings
}

var SSConfig SimpleServerConfig = SimpleServerConfig{} //simple server settings

var XSCR bool = false //server config parser result

func XServerConfigParser(buff []byte) {
	x := string(buff)
	l4g.Info(fmt.Sprintf("get xserver config buff is %s ", x))
	err := json.Unmarshal(buff, &XSS)
	l4g.Info(fmt.Sprintf("get xserver config is %v %v ", XSS, err))
	if err != nil {
		XSCR = false
	} else {
		XSCR = true
		configureLog(&XSS.LogSetting)
	}
}
