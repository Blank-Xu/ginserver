package models

import (
	"ginserver/modules/db"
	"ginserver/modules/log"
)

type Log struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int64 `xorm:"pk autoincr"`
	Level     log.Level
	Type      log.Type
	UserId    int `xorm:"index"`
	RoleId    int `xorm:"index"`
	Method    string
	Path      string
	Params    string
	Ip        string
	Remark    string
	Created   db.JSONTime `xorm:"created"`
}

func (p *Log) TableName() string {
	return "log"
}
