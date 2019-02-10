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
