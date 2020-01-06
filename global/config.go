package global

import (
	"ginserver/pkg/db"
	"ginserver/pkg/log"
	"ginserver/tools/fix"
	"ginserver/tools/httputil"
	"ginserver/tools/redis"
	"ginserver/tools/session"
)

var DefaultConfig *Config

type Config struct {
	RunMode         string           `yaml:"RunMode"`
	AppName         string           `yaml:"AppName"`
	StaticDir       string           `yaml:"StaticDir"`
	TemplateDir     string           `yaml:"TemplateDir"`
	CasbinModelFile string           `yaml:"CasbinModelFile"`
	Log             *log.Option      `yaml:"Log"`
	Session         *session.Option  `yaml:"Session"`
	HttpServer      *httputil.Option `yaml:"HttpServer"`
	Fix             *fix.Fix         `yaml:"Fix"`
	DataBase        []*db.Option     `yaml:"DataBase"`
	Redis           []*redis.Client  `yaml:"Redis"`
}
