package model

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
)

type XServerSettings struct {
	SiteURL                         *string
	ListenAddress                   string
	ConnectionSecurity              *string
	TLSCertFile                     *string
	TLSKeyFile                      *string
	UseLetsEncrypt                  *bool
	LetsEncryptCertificateCacheFile *string
	Forward80To443                  *bool
	ReadTimeout                     *int
	WriteTimeout                    *int
	MaximumLoginAttempts            int
	ServerLocale                    *string
	AllowCorsFrom                   *string
}

type RateLimitSettings struct {
	Enable           *bool
	PerSec           int
	MaxBurst         *int
	MemoryStoreSize  int
	VaryByRemoteAddr bool
	VaryByHeader     string
}

type ServerSettings struct {
	XServerSetting   XServerSettings
	LogSetting       LogSettings
	RateLimitSetting RateLimitSettings
}

var XSS ServerSettings = ServerSettings{} //xserver settings
var XSCR bool = false                     //server config parser result

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
