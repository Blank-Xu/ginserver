package log

import (
	"ginserver/tools/db"
)

type Log struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int64 `xorm:"pk autoincr"`
	Level     Level
	Type      Type
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

func (p *Log) InsertOne() (err error) {
	_, err = db.GetDefaultEngine().InsertOne(p)
	return
}
