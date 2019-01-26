package models

import (
	"time"

	"ginserver/modules/db"
)

type SParams struct {
	*db.Model  `xorm:"-"`
	Id         int    `xorm:"pk autoincr"`
	Name       string `xorm:"unique"`
	Param      string
	Remark     string
	UpdateTime time.Time
}

func (p *SParams) TableName() string {
	return "s_params"
}
