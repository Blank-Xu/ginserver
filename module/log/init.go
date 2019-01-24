package log

import (
	"fmt"
	"io"
	"os"

	"ginserver/module/config"
	"ginserver/module/util"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	var (
		cfgLog  = config.GetConfig().Log
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
	logName := cfgLog.Path + cfgLog.FileName
	logFile, err = os.OpenFile(logName, os.O_WRONLY|os.O_APPEND, 0666)
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
}

func GetLog() *logrus.Logger {
	return log
}
