package global

import (
	"github.com/spf13/viper"
)

var (
	AppName string = "server"
	RunMode string = "debug"
)

var (
	Viper *viper.Viper
)
