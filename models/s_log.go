package models

import (
	"ginserver/modules/db"
)

type LogType int

const (
	LogTypeLogin LogType = iota + 1
	LogTypeLogout
)

type SLog struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int64 `xorm:"pk autoincr"`
	LogType   LogType
	UserId    int `xorm:"index"`
	RoleId    int `xorm:"index"`
	Method    string
	Path      string
	Params    string
	Ip        string
	Remark    string
	Created   db.JSONTime `xorm:"created"`
}

func (p *SLog) TableName() string {
	return "s_log"
}
