package models

import (
	"ginserver/modules/db"
)

type MenuType int

const (
	MenuTypeMain MenuType = iota
	MenuTypeButton
	MenuTypeHref
)

type SMenu struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int `xorm:"pk autoincr"`
	Type      int
	Name      string `xorm:"unique"`
	Method    string
	Path      string
	Icon      string
	Level     int
	OrderNo   int
	State     bool
	ParentId  int
	Created   db.JSONTime `xorm:"created"`
	Updater   int
	Updated   db.JSONTime `xorm:"updated"`
}

func (p *SMenu) TableName() string {
	return "s_menu"
}

func (p *SMenu) SelectByParentId(parentId int) (records []*SMenu, err error) {
	err = db.GetDefaultEngine().SQL(`SELECT *
FROM s_menu
WHERE parent_id = ?
  AND type = ?
  AND state = 1`, parentId, MenuTypeMain).Find(&records)
	return
}

type SRoleMenus struct {
	Id       int           `xorm:"pk autoincr" json:"id"`
	ParentId int           `json:"parent_id"`
	Name     string        `json:"name"`
	Method   string        `json:"method"`
	Path     string        `json:"path"`
	Icon     string        `json:"icon"`
	OrderNo  int           `json:"order_no"`
	List     []*SRoleMenus `xorm:"-" json:"list,omitempty"`
}

func (p *SRoleMenus) SelectMainMenusByUserId(userId int) (records []*SRoleMenus, err error) {
	err = db.GetDefaultEngine().SQL(`SELECT menu.id,
       menu.name,
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
ORDER BY menu.parent_id`, userId, MenuTypeMain).Find(&records)
	return
}

func (p *SRoleMenus) SelectMainMenusByRoleId(roleId int) (records []*SRoleMenus, err error) {
	err = db.GetDefaultEngine().SQL(`SELECT menu.id,
       menu.name,
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
ORDER BY menu.parent_id`, roleId, MenuTypeMain).Find(&records)
	return
}
