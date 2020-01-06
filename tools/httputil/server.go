package httputil

import (
	"fmt"
	"net/http"
	"time"
)

type Option struct {
	IP                  string `yaml:"IP"`
	Port                int    `yaml:"Port"`
	ReadTimeout         int    `yaml:"ReadTimeout"`
	WriteTimeout        int    `yaml:"WriteTimeout"`
	IdleTimeout         int    `yaml:"IdleTimeout"`
	MaxConnsPerHost     int    `yaml:"MaxConnsPerHost"`     // 每一个host对应的最大连接数
	MaxIdleConns        int    `yaml:"MaxIdleConns"`        // 所有host对应的idle状态最大的连接总数
	MaxIdleConnsPerHost int    `yaml:"MaxIdleConnsPerHost"` // 每一个host对应idle状态的最大的连接数
	MaxHeaderBytes      int    `yaml:"MaxHeaderBytes"`
}

func (p *Option) Init() {
	http.DefaultTransport.(*http.Transport).MaxConnsPerHost = p.MaxConnsPerHost
	http.DefaultTransport.(*http.Transport).MaxIdleConns = p.MaxIdleConns
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = p.MaxIdleConnsPerHost
}

func (p *Option) Addr() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func (p *Option) New(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           p.Addr(),
		Handler:        handler,
		ReadTimeout:    time.Second * time.Duration(p.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(p.WriteTimeout),
		IdleTimeout:    time.Second * time.Duration(p.IdleTimeout),
		MaxHeaderBytes: p.MaxHeaderBytes,
	}
}
