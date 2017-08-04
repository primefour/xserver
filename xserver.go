package main

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
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

var xserver_apps = map[string]WebAppIntf{}

func initServer() {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%v", r)
			l4g.Error(err)
		}
	}()
	utils.InitLogSystem()
	//init locale
	utils.InitTranslations()
	//init html templates
	utils.InitHTML()
}

func runApps() {
	for appName, appIntf := range xserver_apps {
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
	for appName, appIntf := range xserver_apps {
		l4g.Info("stop service of %s ", appName)
		appIntf.StopServer()
	}
}

//init locale and log system before start server
func runServer() {
	runApps()
	// wait for kill signal before attempting to gracefully shutdown
	// the running service
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c
	stopApps()
}

func main() {
	initServer()
	runServer()
}
