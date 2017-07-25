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

const (
	L4G_DEBUG_LEVEL = l4g.DEBUG
	MODE_DEV        = "dev"
	MODE_BETA       = "beta"
	MODE_PROD       = "prod"
	LOG_ROTATE_SIZE = 10000
	LOG_FILENAME    = "xserver.log"
)

type ServiceSettings struct {
	SiteURL                                  *string
	ListenAddress                            string
	ConnectionSecurity                       *string
	TLSCertFile                              *string
	TLSKeyFile                               *string
	UseLetsEncrypt                           *bool
	LetsEncryptCertificateCacheFile          *string
	Forward80To443                           *bool
	ReadTimeout                              *int
	WriteTimeout                             *int
	MaximumLoginAttempts                     int
	GoogleDeveloperKey                       string
	EnableOAuthServiceProvider               bool
	EnableIncomingWebhooks                   bool
	EnableOutgoingWebhooks                   bool
	EnableLinkPreviews                       *bool
	AllowCorsFrom                            *string
	SessionLengthWebInDays                   *int
	SessionLengthMobileInDays                *int
	SessionLengthSSOInDays                   *int
	SessionCacheInMinutes                    *int
	WebsocketSecurePort                      *int
	WebsocketPort                            *int
	WebserverMode                            *string
	TimeBetweenUserTypingUpdatesMilliseconds *int64
	ClusterLogTimeoutMilliseconds            *int
	LocalePath                               *string
	ServerLocale                             *string
}

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
	xconfigOK      bool
	apps           map[string]WebAppIntf
	ss             ServiceSettings
	tDefault       i18n.TranslateFunc //system
}

func xserverConfigParser(buff []byte) {
	x := string(buff)
	l4g.Info(fmt.Sprintf("get xserver config buff is %s ", x))
	err := json.Unmarshal(buff, &xserver.ss)
	l4g.Info(fmt.Sprintf("get xserver config is %v %v ", xserver.ss, err))
	if err != nil {
		xserver.xconfigOK = false
	} else {
		xserver.xconfigOK = true
	}
}

type OriginCheckerProc func(*http.Request) bool

func OriginChecker(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return *xserver.ss.AllowCorsFrom == "*" || strings.Contains(*xserver.ss.AllowCorsFrom, origin)
}

var xserver = XServer{
	configFilePath: "./xserver.json",
	xconfig:        utils.NewXConfig("xserver", "./", true, xserverConfigParser),
	xconfigOK:      false,
	apps:           map[string]WebAppIntf{},
	ss:             ServiceSettings{},
}

func GetOriginChecker(r *http.Request) OriginCheckerProc {
	if len(*xserver.ss.AllowCorsFrom) > 0 {
		return OriginChecker
	}
	return nil
}

//just for static add
func AddWebApp(app WebAppIntf) {
	if app == nil {
		return
	}
	appName = app.GetAppName()
	_, ok := xserver.apps[appName]
	if !ok {
		xserver.apps[appName] = app
		//launch app
	} else {
		//only update config
	}
}

func initServer() {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Sprintf("%v", r)
		}
	}()
	//init locale
	utils.InitTranslationsWithDir("i18n")
	//init html templates
	utils.InitHTMLWithDir("templates")
	//load config
	xserver.xconfig.UpdateForce()
	if !xserver.xconfigOK {
		l4g.Error("xserver load config file fail ")
		return
	}
	xserver.tDefault = utils.GetUserTranslations(utils.DEFAULT_LOCALE)

}

func (self *XServer) initLocale() {
	if self.settings.LocalePath == "" {
		self.settings.LocalePath = "i18n"
	}
	self.locales = utils.InitTranslationsWithDir(self.settings.LocalePath)

	if self.locales == nil {
		panic(L4g.Error("locales directory is empty"))
	}

}

func DisableDebugLogForTest() {
	cfgMutex.Lock()
	defer cfgMutex.Unlock()
	if l4g.Global["stdout"] != nil {
		originalDisableDebugLvl = l4g.Global["stdout"].Level
		l4g.Global["stdout"].Level = l4g.ERROR
	}
}

func EnableDebugLogForTest() {
	cfgMutex.Lock()
	defer cfgMutex.Unlock()
	if l4g.Global["stdout"] != nil {
		l4g.Global["stdout"].Level = originalDisableDebugLvl
	}
}

func ConfigureCmdLineLog() {
	ls := model.LogSettings{}
	ls.EnableConsole = true
	ls.ConsoleLevel = "WARN"
	configureLog(&ls)
}

func configureLog(s *model.LogSettings) {

	l4g.Close()

	if s.EnableConsole {
		level := l4g.DEBUG
		if s.ConsoleLevel == "INFO" {
			level = l4g.INFO
		} else if s.ConsoleLevel == "WARN" {
			level = l4g.WARNING
		} else if s.ConsoleLevel == "ERROR" {
			level = l4g.ERROR
		}

		lw := l4g.NewConsoleLogWriter()
		lw.SetFormat("[%D %T] [%L] %M")
		l4g.AddFilter("stdout", level, lw)
	}

	if s.EnableFile {

		var fileFormat = s.FileFormat

		if fileFormat == "" {
			fileFormat = "[%D %T] [%L] %M"
		}

		level := l4g.DEBUG
		if s.FileLevel == "INFO" {
			level = l4g.INFO
		} else if s.FileLevel == "WARN" {
			level = l4g.WARNING
		} else if s.FileLevel == "ERROR" {
			level = l4g.ERROR
		}

		flw := l4g.NewFileLogWriter(GetLogFileLocation(s.FileLocation), false)
		flw.SetFormat(fileFormat)
		flw.SetRotate(true)
		flw.SetRotateLines(LOG_ROTATE_SIZE)
		l4g.AddFilter("file", level, flw)
	}
}

func GetLogFileLocation(fileLocation string) string {
	if fileLocation == "" {
		return FindDir("logs") + LOG_FILENAME
	} else {
		return fileLocation + LOG_FILENAME
	}
}

//init locale and log system before start server
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
