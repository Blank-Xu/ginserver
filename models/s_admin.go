package models

import (
	"time"

	"ginserver/modules/db"
)

type SAdmin struct {
	*db.Model  `xorm:"-"`
	Id         int    `xorm:"pk autoincr"`
	Username   string `xorm:"unique"`
	Password   string
	Salt       string
	Nickname   string
	Email      string
	Phone      string
	Remark     string
	CreatTime  time.Time
	UpdateTime time.Time
	LoginTime  time.Time
	LoginIp    string
}

func (p *SAdmin) TableName() string {
	return "s_admin"
}
