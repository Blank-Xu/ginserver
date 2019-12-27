package s_role

import (
	"strconv"

	"ginserver/models/s_menu"
	"ginserver/pkg/casbin"
	"ginserver/pkg/db"
)

type RoleMenu struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int `xorm:"pk autoincr" json:"id"`
	RoleId    int
	MenuId    int
}

func (p *RoleMenu) TableName() string {
	return "s_role_menu"
}

func (p *RoleMenu) InsertOne() error {
	_, err := db.GetDefaultDB().InsertOne(p)
	// TODO:  casbin.GetEnforcer().AddPermissionForUser(strconv.Itoa(p.RoleId),...)
	return err
}

func (p *RoleMenu) Delete() error {
	_, err := db.GetDefaultDB().Delete(p)
	casbin.GetEnforcer().DeletePermissionsForUser(strconv.Itoa(p.RoleId))
	return err
}

type RoleMenuDetail struct {
	Id          int               `xorm:"pk autoincr" json:"id"`
	ParentId    int               `json:"parent_id"`
	SubIds      map[int]bool      `xorm:"-",json:"-"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Icon        string            `json:"icon"`
	OrderNo     int               `json:"order_no"`
	List        []*RoleMenuDetail `xorm:"-" json:"list,omitempty"`
}

func (p *RoleMenuDetail) SelectMainMenuByUserId(userId int) ([]*RoleMenuDetail, error) {
	var records []*RoleMenuDetail
	return records, db.GetDefaultDB().SQL(`SELECT menu.id,
       menu.name,
       menu.description,
       menu.method,
       menu.path,
       menu.icon,
       menu.order_no,
       menu.parent_id
FROM s_user user
       LEFT JOIN s_user_role user_role ON user.id = user_role.user_id
       LEFT JOIN s_role_menu role_menu ON role_menu.role_id = user_role.role_id
       LEFT JOIN s_menu menu ON menu.id = role_menu.menu_id
WHERE user.id = ?
  AND menu.type = 0
  AND menu.state = 1
ORDER BY menu.parent_id`, userId, s_menu.TypeMain).Find(&records)
}

func (p *RoleMenuDetail) SelectMainMenuByRoleId(roleId int) ([]*RoleMenuDetail, error) {
	var records []*RoleMenuDetail
	return records, db.GetDefaultDB().SQL(`SELECT menu.id,
       menu.name,
       menu.description,
       menu.method,
       menu.path,
       menu.icon,
       menu.order_no,
       menu.parent_id
FROM s_role_menu role_menu
       LEFT JOIN s_menu menu ON menu.id = role_menu.menu_id
WHERE role_menu.role_id = ?
  AND menu.type = ?
  AND menu.state = 1
ORDER BY menu.parent_id`, roleId, s_menu.TypeMain).Find(&records)
}
