package casbin

import (
	"ginserver/tools/db"
)

type casbinRoleMenu struct {
	RoleId string
	Path   string
	Method string
}

func loadRoleMenuPolicy() error {
	var rules []*casbinRoleMenu
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
		defaultEnforcer.AddPermissionForUser(rule.RoleId, rule.Path, rule.Method)
	}
	return nil
}

type casbinUserRole struct {
	UserId string
	RoleId string
}

func loadUserRolePolicy() error {
	var rules []*casbinUserRole
	if err := db.GetDefaultEngine().SQL("SELECT * FROM s_user_role").Find(&rules); err != nil {
		return err
	}
	for _, rule := range rules {
		defaultEnforcer.AddRoleForUser(rule.UserId, rule.RoleId)
	}
	return nil
}
