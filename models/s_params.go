package models

import (
	"ginserver/modules/db"
	"ginserver/modules/utils"
)

type SParams struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int            `xorm:"pk autoincr" json:"id"`
	Name      string         `xorm:"unique"`
	Param     string         `json:"param"`
	Remark    string         `json:"remark"`
	Updated   utils.JSONTime `json:"updated"`
}

func (p *SParams) TableName() string {
	return "s_params"
}
