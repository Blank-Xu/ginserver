package init

import (
	"fmt"
	"io"
	"os"
	"strings"

	"ginserver/init/config"
	"ginserver/tools/utils"

	"github.com/sirupsen/logrus"
)

func logInit() {
	defaultLog := logrus.StandardLogger()

	var (
		cfg     = config.GetConfig().Log
		logName string
		logFile *os.File
		err     error
	)

	// check log dir
	if !utils.IsFileExit(cfg.Path) {
		if err = os.Mkdir(cfg.Path, 0777); err != nil {
			panic(fmt.Sprintf("Create log path [%s] err: [%v]", cfg.Path, err))
		}
	}

	// TODO: need to split log file
	// open or create log file
	if strings.HasSuffix(cfg.Path, "/") {
		logName = cfg.Path + cfg.FileName
	} else {
		logName = cfg.Path + "/" + cfg.FileName
	}
	logFile, err = os.OpenFile(logName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Sprintf("Open or Create log file [%s] err: [%v]", logName, err))
	}

	format := &logrus.JSONFormatter{
		TimestampFormat: utils.TimeLayoutDefault,
	}
	if config.GetConfig().RunMode == "debug" {
		defaultLog.SetOutput(io.MultiWriter(logFile, os.Stdout))
		format.PrettyPrint = true
	} else {
		defaultLog.SetOutput(logFile)
	}
	defaultLog.SetFormatter(format)
	defaultLog.SetLevel(logrus.Level(cfg.Level))
	defaultLog.SetReportCaller(cfg.ReportCaller)
}
