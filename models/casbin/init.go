package casbin

import (
	"fmt"
	"ginserver/modules/config"

	"github.com/casbin/casbin"
)

var (
	enforcer *casbin.Enforcer
)

func Init() {
	var cfg = config.GetConfig()
	enforcer = casbin.NewEnforcer(cfg.CasbinFile, false)
	if err := loadRoleMenuPolicy(); err != nil {
		panic(fmt.Sprintf("load casbin role menu policy error, err: [%v]", err))
	}
	if err := loadUserRolePolicy(); err != nil {
		panic(fmt.Sprintf("load casbin user role policy error, err: [%v]", err))
	}
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
