package init

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"ginserver/init/config"
)

var configFile = flag.String("config", "configs/app_debug.yaml", "config file")

func Init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	fmt.Printf("Server Starting ... \n - version: [%s]  \n - args: %s\n", config.Version, os.Args)
	fmt.Printf("Read Config File ... \n - file_name: [%s]\n", *configFile)
	fmt.Println(" - you can use [-config file] command to set config file when server start.")

	config.Init(*configFile)
	fmt.Println("Read Config Success")
	// fix default setting
	fix()
	// start log first
	logInit()
	logrus.Info("Server Starting ...")

	redisInit()

	dbsInit()

	modelsInit()

	casbinInit()
}
