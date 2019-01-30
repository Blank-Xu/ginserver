package log

import (
	"fmt"
	"ginserver/modules/utils"
	"io"
	"os"
	"strings"

	"ginserver/modules/config"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func Init() {
	log = logrus.StandardLogger()

	var (
		cfg     = config.GetConfig().Log
		logName string
		logFile *os.File
		err     error
	)

	// check log dir
	if !utils.IsFileExit(cfg.Path) {
		if err = os.Mkdir(cfg.Path, 0777); err != nil {
			panic(fmt.Sprintf("create log path [%s] err: [%v]", cfg.Path, err))
		}
	}

	// open or create log file
	if strings.HasSuffix(cfg.Path, "/") {
		logName = cfg.Path + cfg.FileName
	} else {
		logName = cfg.Path + "/" + cfg.FileName
	}
	logFile, err = os.OpenFile(logName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Sprintf("open or create log name [%s] err: [%v]", logName, err))
	}

	format := &logrus.JSONFormatter{
		TimestampFormat: utils.TimeLayoutDefault,
	}
	if config.GetConfig().RunMode == "debug" {
		log.SetOutput(io.MultiWriter(logFile, os.Stdout))
		format.PrettyPrint = true
	} else {
		log.SetOutput(logFile)
	}
	log.SetFormatter(format)
	log.SetLevel(logrus.Level(cfg.Level))
	log.SetReportCaller(cfg.ReportCaller)
}
