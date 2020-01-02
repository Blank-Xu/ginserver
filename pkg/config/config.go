package config

import (
	"ginserver/tools/fix"
	"ginserver/tools/httputil"
	"ginserver/tools/session"
)

var Default *Config

type Config struct {
	RunMode         string           `yaml:"RunMode"`
	AppName         string           `yaml:"AppName"`
	StaticDir       string           `yaml:"StaticDir"`
	TemplateDir     string           `yaml:"TemplateDir"`
	CasbinModelFile string           `yaml:"CasbinModelFile"`
	Session         *session.Session `yaml:"Session"`
	HttpServer      *httputil.Server `yaml:"HttpServer"`
	Fix             *fix.Fix         `yaml:"Fix"`
	// Log         *Log        `yaml:"Log"`
	// DataBase    []*Database `yaml:"DataBase"`
	// Redis       *redis      `yaml:"Redis"`
}

func (p *Config) Init() error {

	return nil
}
