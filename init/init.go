package init

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"ginserver/global"
)

var (
	configFile = flag.String("config", "configs/app_debug.yaml", "config file")
)

const (
	ConfigDir  = "configs"
	ServerName = "webserver"
	Version    = "0.1.0"
)

const (
	remoteConfigRefreshTime = time.Second * 5
)

var (
	runMode = flag.String("m", "debug", "run mode")

	// for remote config setting
	remote        = flag.Bool("r", false, "remote")
	provider      = flag.String("P", "etcd", "provider")
	endpoint      = flag.String("h", "http://127.0.0.1:4001", "endpoint")
	wpath         = flag.String("p", ServerName, "path")
	secretkeyring = flag.String("s", ServerName, "secretkeyring")
	configType    = flag.String("t", "yaml", "configType")
)

// type config struct {
// 	Fix    *fix.Fix               `json:"fix" yaml:"fix"`
// 	Server *httpserver.HttpServer `json:"server" yaml:"server"`
// 	Jwt    []*jwt.Jwt             `json:"jwt" yaml:"jwt"`
// 	Log    *wlog.Log              `json:"log" yaml:"log"`
// }

// func Init() {
// 	flag.Parse()
// 	log.Printf("server version: %s, start args: %v\n", Version, flag.Args())
//
// 	if *remote {
// 		parseRemoteConfig(loadConfig)
// 	} else {
// 		parseLocalConfig(loadConfig)
// 	}
//
// 	log.SetOutput(io.MultiWriter(os.Stdout))
//
// 	if err := casbin.Init(config.Default.CasbinModelFile); err != nil {
// 		panic(err)
// 	}
// }

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

	if err = yaml.Unmarshal(byt, nil); err != nil {
		panic(fmt.Sprintf("read config error, file: [%s], err: [%v]", filename, err))
	}

}
