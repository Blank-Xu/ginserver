package session

import (
	"errors"
	"fmt"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
)

type Session struct {
	Provider string `yaml:"Provider"`
	Path     string `yaml:"Path"`
	Domain   string `yaml:"Domain"`
	Secret   string `yaml:"Secret"`
	MaxAge   int    `yaml:"MaxAge"`
	HttpOnly bool   `yaml:"HttpOnly"`
	RedisPassword string `yaml:"redis_password"`
}

func (p *Session) NewStore() (store sessions.Store,err error) {
	switch p.Provider {
	case "redis":
		store, err = redis.NewStore(30, "tcp", p.Domain, p.RedisPassword, []byte(p.Secret))
		if err != nil {
			return nil,  fmt.Errorf("create redis session failed, err: %v", err)
		}
	case "memstore":
		store = memstore.NewStore([]byte(p.Secret))
	default:
		return nil, errors.New("not support this session")
	}
	
	if len(p.Path) ==0{
		p.Path = "/"
	}
	
	store.Options(sessions.Options{
		Path:     p.Path,
		MaxAge:   p.MaxAge,
		HttpOnly: p.HttpOnly,
	})
	
	return
}