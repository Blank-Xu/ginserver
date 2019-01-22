package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg = new(config)

func Init(file string) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("load config error, file: [%s], err: [%v]", file, err))
	}

	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Sprintf("read config error, file: [%s], err: [%v]", file, err))
	}
}

func GetConfig() *config {
	return cfg
}
