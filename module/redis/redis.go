package redis

import "ginserver/module/config"

func Init() {
	var cfg = config.GetConfig().Redis
}
