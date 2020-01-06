package log

import (
	"github.com/rs/zerolog"

	"ginserver/pkg/db"
	"ginserver/tools/timeutil"
)

type Log struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int64 `xorm:"pk autoincr"`
	Level     zerolog.Level
	Type      Type
	UserId    int `xorm:"index"`
	RoleId    int `xorm:"index"`
	Method    string
	Path      string
	Params    string
	Ip        string
	Remark    string
	Created   timeutil.JSONTime `xorm:"created"`
}

func (p *Log) TableName() string {
	return "log"
}

func (p *Log) InsertOne() (err error) {
	_, err = db.GetDefaultDB().InsertOne(p)
	return
}

func Insert(level zerolog.Level, lType Type, userId, roleId int, method, path, params, ip, remark string) (err error) {
	_, err = db.GetDefaultDB().Exec(
		"INSERT INTO log (level, type, user_id, role_id, method, path, params, ip, remark) VALUES (?,?,?,?,?,?,?,?,?)", level, lType, userId, roleId, method, path, params, ip, remark)
	return
}
