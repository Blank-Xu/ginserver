package init

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"ginserver/global"
	"ginserver/pkg/casbin"
	"ginserver/pkg/db"
)

var (
	// for local config
	configFile = flag.String("c", "configs/app_debug.yaml", "config file")

	// for remote config
	remote        = flag.Bool("r", false, "remote")
	configType    = flag.String("t", "yaml", "configType")
	provider      = flag.String("P", "etcd", "provider")
	endpoint      = flag.String("h", "http://127.0.0.1:4001", "endpoint")
	path          = flag.String("p", "path", "path")
	secretkeyring = flag.String("s", "secretkeyring", "secretkeyring")
)

func Init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	log.Printf("server starting ...\n -- version: [%s]\n -- args: %v", global.Version, flag.Args())

	defaultViper := viper.GetViper()
	var err error
	if *remote {
		err = parseRemoteConfig(defaultViper, *configType, *provider, *endpoint, *path, *secretkeyring, loadConfig)
	} else {
		err = parseLocalConfig(defaultViper, *configFile, loadConfig)
	}
	if err != nil {
		panic(fmt.Sprintf("parse config failed, err: %v", err))
	}

	global.Viper = defaultViper

	var cfg global.Config
	if err = defaultViper.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("Unmarshal config failed, err: %v", err))
	}
	global.DefaultConfig = &cfg
	global.AppName = cfg.AppName
	global.RunMode = cfg.RunMode

	cfg.Fix.Init()

	if err = db.SetDBS(cfg.DataBase); err != nil {
		panic(err)
	}

	if err = casbin.Init(cfg.CasbinModelFile); err != nil {
		panic(fmt.Sprintf("load casbin failed, err: %v", err))
	}

	cfg.HttpServer.Init()

	// log.SetOutput(io.MultiWriter(os.Stdout))
}
