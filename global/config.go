package global

import (
	"ginserver/pkg/db"
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
	Session         *session.Session `yaml:"Session"`
	HttpServer      *httputil.Server `yaml:"HttpServer"`
	Fix             *fix.Fix         `yaml:"Fix"`
	DataBase        []*db.Option     `yaml:"DataBase"`
	Redis           []*redis.Client  `yaml:"Redis"`
	// Log         *Log        `yaml:"Log"`
}
