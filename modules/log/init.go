package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"ginserver/modules/config"
	"ginserver/modules/util"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	var (
		cfgLog  = config.GetConfig().Log
		logName string
		logFile *os.File
		err     error
	)

	// check log dir
	if !util.IsFileExit(cfgLog.Path) {
		if err = os.Mkdir(cfgLog.Path, 0777); err != nil {
			panic(fmt.Sprintf("create log path [%s] err: [%v]", cfgLog.Path, err))
		}
	}

	// open or create log file
	if strings.HasSuffix(cfgLog.Path, "/") {
		logName = cfgLog.Path + cfgLog.FileName
	} else {
		logName = cfgLog.Path + "/" + cfgLog.FileName
	}
	logFile, err = os.OpenFile(logName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Sprintf("open or create log name [%s] err: [%v]", logName, err))
	}

	format := &logrus.JSONFormatter{
		TimestampFormat: util.TimeLayoutDefault,
	}
	if config.GetConfig().RunMode == "debug" {
		log.SetOutput(io.MultiWriter(logFile, os.Stdout))
		format.PrettyPrint = true
	} else {
		log.SetOutput(logFile)
	}
	log.SetFormatter(format)
	log.SetLevel(logrus.Level(cfgLog.Level))
	log.SetReportCaller(true)
}

func GetLog() *logrus.Logger {
	return log
}
