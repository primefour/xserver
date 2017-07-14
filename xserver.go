package main

import (
	"flag"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/api"
	"github.com/primefour/xserver/app"
	"github.com/primefour/xserver/utils"
	"os"
	"os/signal"
	"syscall"
)

func doLoadConfig(fileName string) (err string) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Sprintf("%v", r)
		}
	}()
	utils.TranslationsPreInit()
	utils.EnableConfigFromEnviromentVars()
	utils.LoadConfig(fileName)
	utils.InitializeConfigWatch()
	utils.EnableConfigWatch()

	return ""
}

func runServer(configFile string) {
	l4g.Debug("configure file path is %s ", configFile)
	if errstr := doLoadConfig(configFile); errstr != "" {
		l4g.Exit("Unable to load mattermost configuration file: ", errstr)
		return
	}

	utils.InitTranslations(utils.Cfg.LocalizationSettings)

	//pwd, _ := os.Getwd()

	app.NewServer()
	app.InitStores()
	api.InitRouter()
	api.InitApi(false)

	if len(utils.Cfg.SqlSettings.DataSourceReplicas) > 1 {
		l4g.Warn(utils.T("store.sql.read_replicas_not_licensed.critical"))
		utils.Cfg.SqlSettings.DataSourceReplicas = utils.Cfg.SqlSettings.DataSourceReplicas[:1]
	}

	app.StartServer()

	setDiagnosticId()
	utils.RegenerateClientConfig()
	go runSecurityJob()
	go runDiagnosticsJob()

	go runTokenCleanupJob()

	// wait for kill signal before attempting to gracefully shutdown
	// the running service
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	app.StopServer()
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

var config_dir string

func init() {
	flag.StringVar(&config_dir, "config", "./config", "config file for server")
}

func main() {
	flag.Parse()
	fileName := config_dir + "/config.json"
	runServer(fileName)
}
