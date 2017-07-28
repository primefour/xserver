package model

import (
	"github.com/primefour/xserver/utils"
)

const (
	LOG_ROTATE_SIZE = 10000
	LOG_FILENAME    = "xserver.log"
	LOG_DIRNAME     = "sklog"
)

type LogSettings struct {
	EnableConsole          bool
	ConsoleLevel           string
	EnableFile             bool
	FileLevel              string
	FileFormat             string
	FileLocation           string
	EnableWebhookDebugging bool
	EnableDiagnostics      *bool
}

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
}

type ServiceSettings struct {
	XServerSetting XServerSettings
	LogSetting     LogSettings
}

var XServiceSetting ServiceSettings = ServiceSettings{}
var XServerConfigResult string = false

func XServerConfigParser(buff []byte) {
	x := string(buff)
	l4g.Info(fmt.Sprintf("get xserver config buff is %s ", x))
	err := json.Unmarshal(buff, &XServiceSettings)
	l4g.Info(fmt.Sprintf("get xserver config is %v %v ", XServiceSettings, err))
	if err != nil {
		XServerConfigResult = false
	} else {
		XServerConfigResult = true
		configureLog(XServiceSettings.LogSetting)
	}
}

func DisableDebugLogForTest() {
	if l4g.Global["stdout"] != nil {
		l4g.Global["stdout"].Level = l4g.ERROR
	}
}

func EnableDebugLogForTest() {
	if l4g.Global["stdout"] != nil {
		l4g.Global["stdout"].Level = l4g.DEBUG
	}
}

func ConfigureCmdLineLog() {
	ls := LogSettings{}
	ls.EnableConsole = true
	ls.ConsoleLevel = "WARN"
	configureLog(&ls)
}

func configureLog(s *LogSettings) {
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

		flw := l4g.NewFileLogWriter(getLogFileLocation(s.FileLocation), false)
		flw.SetFormat(fileFormat)
		flw.SetRotate(true)
		flw.SetRotateLines(LOG_ROTATE_SIZE)
		l4g.AddFilter("file", level, flw)
	}
}

func getLogFileLocation(fileLocation string) string {
	if fileLocation == "" {
		return utils.FindDir(LOG_DIRNAME) + LOG_FILENAME
	} else {
		return fileLocation + LOG_FILENAME
	}
}
