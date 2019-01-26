package models

import (
	"time"

	"ginserver/modules/db"
)

type SLog struct {
	*db.Model  `xorm:"-"`
	Id         int64 `xorm:"pk autoincr"`
	Uid        int   `xorm:"index"`
	Role       string
	Method     string
	Router     string
	Params     string
	CreateTime time.Time
	Ip         string
	Remark     string
}

func (p *SLog) TableName() string {
	return "s_log"
}
