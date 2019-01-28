package models

import (
	"ginserver/modules/db"
	"ginserver/modules/util"
)

type SAdmin struct {
	*db.Model  `xorm:"-" json:"-"`
	Id         int64         `xorm:"pk autoincr" json:"id"`
	Username   string        `xorm:"unique" json:"username"`
	Password   string        `json:"-"`
	Salt       string        `json:"-"`
	Nickname   string        `json:"nickname"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	Remark     string        `json:"remark"`
	RegisterIp string        `json:"-"`
	Created    util.JSONTime `xorm:"created" json:"created"`
	Updated    util.JSONTime `xorm:"updated" json:"updated"`
	LoginTime  util.JSONTime `json:"login_time"`
	LoginIp    string        `json:"login_ip"`
}

func NewSAdmin(id int64) *SAdmin {
	return &SAdmin{Id: id}
}

const (
	s_admin_update_cols = ""
	s_admin_update_omit = "id,"
)

func (p *SAdmin) TableName() string {
	return "s_admin"
}
