package init

import (
	"flag"

	"github.com/rs/zerolog/log"
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

	log.Info().Msgf("server starting, version: [%s], args: %v", global.Version, flag.Args())

	defaultViper := viper.GetViper()
	var err error
	if *remote {
		err = parseRemoteConfig(defaultViper, *configType, *provider, *endpoint, *path, *secretkeyring, loadConfig)
	} else {
		err = parseLocalConfig(defaultViper, *configFile, loadConfig)
	}
	if err != nil {
		log.Panic().Msg("parse config failed")
	}

	global.Viper = defaultViper

	var cfg global.Config
	if err = defaultViper.Unmarshal(&cfg); err != nil {
		log.Panic().Msg("Unmarshal config failed")
	}
	global.DefaultConfig = &cfg
	global.AppName = cfg.AppName
	global.RunMode = cfg.RunMode

	cfg.Fix.Init()

	if err = cfg.Log.Init(); err != nil {
		log.Panic().Err(err).Msgf("zerolog init failed")
	}

	if err = db.SetDBS(cfg.DataBase); err != nil {
		log.Panic().Err(err).Msg("connect database failed")
	}

	if err = casbin.Init(cfg.CasbinModelFile); err != nil {
		log.Panic().Err(err).Msg("casbin load failed")
	}

	cfg.HttpServer.Init()
}
