package init

import (
	"fmt"

	"ginserver/init/config"
	defaultCasbin "ginserver/tools/casbin"
	"ginserver/tools/db"

	"github.com/casbin/casbin/v2"
)

func casbinInit() {
	var cfg = config.GetConfig()
	enforce, err := casbin.NewEnforcer(cfg.RbacModel, false)
	if err != nil {
		panic("create casbin enforcer failed, err: " + err.Error())
	}

	enforcer := &enforcer{enforce}
	// load rules
	if err := enforcer.loadRoleMenuPolicy(); err != nil {
		panic(fmt.Sprintf("Load casbin role menu policy failed, err: [%v]", err))
	}
	if err := enforcer.loadUserRolePolicy(); err != nil {
		panic(fmt.Sprintf("Load casbin user role policy failed, err: [%v]", err))
	}

	defaultCasbin.SetEnforcer(enforcer.Enforcer)
}

type enforcer struct {
	*casbin.Enforcer
}

func (p *enforcer) loadRoleMenuPolicy() error {
	var rules []*struct {
		RoleId string
		Path   string
		Method string
	}
	err := db.GetDefaultEngine().SQL(
		`SELECT role_menu.role_id,
       menu.path,
       menu.method
FROM s_role_menu role_menu
       LEFT JOIN s_menu menu ON menu.id = role_menu.menu_id
WHERE menu.method IN ('GET', 'POST', 'PUT', 'PATCH', 'HEAD', 'OPTIONS', 'DELETE', 'CONNECT', 'TRACE')
  AND menu.state = 1
ORDER BY menu.parent_id`).Find(&rules)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		p.Enforcer.AddPermissionForUser(rule.RoleId, rule.Path, rule.Method)
	}
	return nil
}

func (p *enforcer) loadUserRolePolicy() error {
	var rules []*struct {
		UserId string
		RoleId string
	}
	if err := db.GetDefaultEngine().SQL("SELECT * FROM s_user_role").Find(&rules); err != nil {
		return err
	}
	for _, rule := range rules {
		p.Enforcer.AddRoleForUser(rule.UserId, rule.RoleId)
	}
	return nil
}
