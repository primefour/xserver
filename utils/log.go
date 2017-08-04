package utils

import (
	"bytes"
	l4g "github.com/alecthomas/log4go"
	"io"
	"io/ioutil"
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

var defaultLogSetting = LogSettings{
	EnableConsole: true,
	ConsoleLevel:  "DEBUG",
	EnableFile:    true,
	FileLevel:     "DEBUG",
}

func InitLogSystem() {
	configureLog(&defaultLogSetting)
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

// DebugReader logs the content of the io.Reader and returns a new io.Reader
// with the same content as the received io.Reader.
// If you pass reader by reference, it won't be re-created unless the loglevel
// includes Debug.
// If an error is returned, the reader is consumed an cannot be read again.
func DebugReader(reader io.Reader, message string) (io.Reader, error) {
	var err error
	l4g.Debug(func() string {
		content, err := ioutil.ReadAll(reader)
		if err != nil {
			return ""
		}

		reader = bytes.NewReader(content)
		return message + string(content)
	})

	return reader, err
}
