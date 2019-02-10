package models

import (
	"ginserver/modules/db"
)

type SUserRole struct {
	*db.Model `xorm:"-" json:"-"`
	UserId    int
	RoleId    int
}

func (p *SUserRole) TableName() string {
	return "s_user_role"
}
