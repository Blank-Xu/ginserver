package models

import (
	"ginserver/modules/db"
)

type SRoleMenu struct {
	*db.Model `xorm:"-" json:"-"`
	RoleId    int
	MenuId    int
}

func (p *SRoleMenu) TableName() string {
	return "s_role_menu"
}

func (p *SRoleMenu) InsertOne() error {
	_, err := db.GetDefaultEngine().InsertOne(p)
	return err
}

type SRoleMenuDetail struct {
	Id       int                `xorm:"pk autoincr" json:"id"`
	ParentId int                `json:"parent_id"`
	Name     string             `json:"name"`
	Method   string             `json:"method"`
	Path     string             `json:"path"`
	Icon     string             `json:"icon"`
	OrderNo  int                `json:"order_no"`
	List     []*SRoleMenuDetail `xorm:"-" json:"list,omitempty"`
}

func (p *SRoleMenuDetail) SelectMainMenuByUserId(userId int) (records []*SRoleMenuDetail, err error) {
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

func (p *SRoleMenuDetail) SelectMainMenuByRoleId(roleId int) (records []*SRoleMenuDetail, err error) {
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
