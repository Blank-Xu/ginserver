package init

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"ginserver/global"
)

var (
	configFile = flag.String("config", "configs/app_debug.yaml", "config file")

	defaultConfig *config
)

func Init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	filename := *configFile

	log.Printf("Server Starting ... \n - version: [%s]  \n - args: %s\n", global.Version, os.Args)
	log.Printf("Read Config File ... \n - filename: [%s]\n", filename)
	log.Println(" - you can use [-config file] command to set config file when server start.")

	byt, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("load config error, file: [%s], err: [%v]", filename, err))
	}

	if err = yaml.Unmarshal(byt, defaultConfig); err != nil {
		panic(fmt.Sprintf("read config error, file: [%s], err: [%v]", filename, err))
	}

}

func GetConfig() *config {
	return defaultConfig
}
