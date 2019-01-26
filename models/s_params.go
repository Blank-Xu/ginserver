package models

import (
	"time"

	"ginserver/modules/db"
)

type SParams struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int       `xorm:"pk autoincr" json:"id"`
	Name      string    `xorm:"unique"`
	Param     string    `json:"param"`
	Remark    string    `json:"remark"`
	Updated   time.Time `json:"updated"`
}

func (p *SParams) TableName() string {
	return "s_params"
}
