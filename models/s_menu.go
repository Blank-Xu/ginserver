package models

import (
	"ginserver/modules/db"
)

type SMenu struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int    `xorm:"pk autoincr"`
	Name      string `xorm:"unique"`
	Path      string
	Icon      string
	Level     int
	OrderNo   int
	State     bool
	Created   db.JSONTime `xorm:"created"`
	Updater   int
	Updated   db.JSONTime `xorm:"updated"`
}

func (p *SMenu) TableName() string {
	return "s_menu"
}
