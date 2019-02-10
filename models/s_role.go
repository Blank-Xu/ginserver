package models

import (
	"ginserver/modules/db"
)

type SRole struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int    `xorm:"pk autoincr" json:"id"`
	Name      string `xorm:"unique"`
	State     bool
	Remark    string
	Created   db.JSONTime `xorm:"created"`
	Updater   int
	Updated   db.JSONTime `xorm:"updated"`
}

func (p *SRole) TableName() string {
	return "s_role"
}

func (p *SRole) SelectOneByUserId(userId int) (bool, error) {
	return db.GetDefaultEngine().SQL(`SELECT role.*
	FROM s_user_role user_role
	LEFT JOIN s_role role ON role.id = user_role.role_id
	WHERE user_role.user_id = ?`, userId).Get(p)
}
