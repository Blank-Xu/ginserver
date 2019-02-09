package models

import (
	"ginserver/modules/db"
)

type LogType int

const (
	LogTypeLogin LogType = iota + 1
	LogTypeLogout
)

type LogLevel int

const (
	// 0:panic, 1:fatal, 2:error, 3:warn, 4:info, 5:debug, 6:trace
	LogLevelPanic LogLevel = iota
	LogLevelFatal
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

type SLog struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int64 `xorm:"pk autoincr"`
	Level     LogLevel
	Type      LogType
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
