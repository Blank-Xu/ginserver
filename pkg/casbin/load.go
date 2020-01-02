package casbin

import (
	"errors"

	"github.com/casbin/casbin/v2"

	"ginserver/pkg/db"
)

func addRoleForUser(enforcer *casbin.Enforcer) error {
	var rules []*struct {
		UserId string
		RoleId string
	}
	if err := db.GetDefaultDB().SQL("SELECT * FROM s_user_role").Find(&rules); err != nil {
		return err
	}

	for _, rule := range rules {
		ok, err := enforcer.AddRoleForUser(rule.UserId, rule.RoleId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("casbin add role for user failed")
		}
	}

	return nil
}

func addAllPermissionForUser(enforcer *casbin.Enforcer) error {
	var rules []*struct {
		RoleId string
		Path   string
		Method string
	}
	err := db.GetDefaultDB().SQL(
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
		ok, err := enforcer.AddPermissionForUser(rule.RoleId, rule.Path, rule.Method)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("casbin add permission for user failed")
		}
	}

	return nil
}
