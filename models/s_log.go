package models

import (
	"ginserver/modules/db"
	"ginserver/modules/utils"
)

type SLog struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int64 `xorm:"pk autoincr"`
	Uid       int   `xorm:"index"`
	Role      string
	Method    string
	Router    string
	Params    string
	Created   utils.JSONTime
	Ip        string
	Remark    string
}

func (p *SLog) TableName() string {
	return "s_log"
}
