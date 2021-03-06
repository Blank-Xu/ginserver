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

func Insert(level Level, lType Type, userId, roleId int, method, path, params, ip, remark string) (err error) {
	_, err = db.GetDefaultEngine().Exec(
		"INSERT INTO log (level, type, user_id, role_id, method, path, params, ip, remark) VALUES (?,?,?,?,?,?,?,?,?)", level, lType, userId, roleId, method, path, params, ip, remark)
	return
}
