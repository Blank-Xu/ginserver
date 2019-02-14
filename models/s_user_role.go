package models

import (
	"strconv"

	"ginserver/modules/casbin"
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

func (p *SUserRole) InsertOne() error {
	_, err := db.GetDefaultEngine().InsertOne(p)
	if err != nil {
		casbin.GetEnforcer().AddRoleForUser(strconv.Itoa(p.UserId), strconv.Itoa(p.RoleId))
	}
	return err
}

func (p *SUserRole) Delete() error {
	_, err := db.GetDefaultEngine().Delete(p)
	if err != nil {
		casbin.GetEnforcer().DeleteUser(strconv.Itoa(p.UserId))
	}
	return err
}
