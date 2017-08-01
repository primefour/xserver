package main

import (
	"flag"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/api"
	"github.com/primefour/xserver/app"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/utils"
	"os"
	"os/signal"
	"syscall"
)

const (
	L4G_DEBUG_LEVEL = l4g.DEBUG
	MODE_DEV        = "dev"
	MODE_BETA       = "beta"
	MODE_PROD       = "prod"
)

type WebAppIntf interface {
	NewInstance()
	InitStores()
	InitRouter()
	InitApi()
	StartServer()
	StopServer()
	LoadConfig() bool
	GetAppName() string
}

type XServer struct {
	configFilePath string
	xconfig        *utils.XConfig
	apps           map[string]WebAppIntf
	tDefault       i18n.TranslateFunc //system
}

type OriginCheckerProc func(*http.Request) bool

func OriginChecker(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return *model.XServiceSetting.AllowCorsFrom == "*" || strings.Contains(*model.XServiceSetting.AllowCorsFrom, origin)
}

func GetOriginChecker(r *http.Request) OriginCheckerProc {
	if len(*model.XServiceSetting.AllowCorsFrom) > 0 {
		return OriginChecker
	}
	return nil
}

func initServer() {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Sprintf("%v", r)
		}
	}()
	xserver.xconfig = utils.NewXConfig("xserver", xserver.configFilePath, true, model.XServerConfigParser)
	//init locale
	utils.InitTranslations()
	//init html templates
	utils.InitHTML()
	//load config
	xserver.xconfig.UpdateForce()

	if !model.XServerConfigResult {
		l4g.Error("xserver load config file fail ")
		return
	}

	//get translate function
	if len(*model.XServiceSetting.ServerLocale) != 0 {
		xserver.tDefault = utils.GetUserTranslations(*model.XServiceSetting.ServerLocale)
	} else {
		xserver.tDefault = utils.GetUserTranslations(utils.DEFAULT_LOCALE)
	}

}

func runApps() {
	for appName, appIntf := range xserver.apps {
		name = appIntf.GetAppName()
		if appName != name {
			l4g.Error("Register Name is not consistent with actual name")
			continue
		}
		if !appIntf.LoadConfig() {
			l4g.Error(fmt.Sprintf("%s load config fail ", appName))
			continue
		}

		if !appIntf.NewInstance() {
			l4g.Error(fmt.Sprintf("%s new Instance fail ", appName))
			continue
		}

		if !appIntf.InitStores() {
			l4g.Error(fmt.Sprintf("%s init stores fail ", appName))
			continue
		}

		if !appIntf.InitRouter() {
			l4g.Error(fmt.Sprintf("%s init route fail ", appName))
			continue
		}
		if !appIntf.InitApi() {
			l4g.Error(fmt.Sprintf("%s init api fail ", appName))
			continue
		}
		if !appIntf.StartServer() {
			l4g.Error(fmt.Sprintf("%s start server fail ", appName))
			continue
		}
	}
}

func stopApps() {
	for appName, appIntf := range xserver.apps {
		l4g.Info("stop service of %s ", appName)
		appIntf.StopServer()
	}
}

//init locale and log system before start server
func runServer() {
	runApps()
	go runSecurityJob()
	go runDiagnosticsJob()
	go runTokenCleanupJob()
	// wait for kill signal before attempting to gracefully shutdown
	// the running service
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c
	stopApps()
}

func runSecurityJob() {
	doSecurity()
}

func runDiagnosticsJob() {
	doDiagnostics()
}

func runTokenCleanupJob() {
	doTokenCleanup()
}

func resetStatuses() {
}

func setDiagnosticId() {
}

func doSecurity() {
}

func doDiagnostics() {
}

func doTokenCleanup() {
	app.Srv.Store.Token().Cleanup()
}

//static add app
var xserver = XServer{
	apps: map[string]WebAppIntf{},
}

func init() {
	flag.StringVar(&xserver.configFilePath, "config", "./config", "config file for server")
}

func main() {
	flag.Parse()
	initServer()
	runServer()
}
