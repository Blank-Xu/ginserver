package models

import (
	"ginserver/modules/db"
	"ginserver/modules/util"
)

type SAdmin struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int           `xorm:"pk autoincr" json:"id"`
	Username  string        `xorm:"unique" json:"username"`
	Password  string        `json:"-"`
	Salt      string        `json:"-"`
	Nickname  string        `json:"nickname"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Remark    string        `json:"remark"`
	Created   util.JSONTime `json:"created"`
	Updated   util.JSONTime `json:"updated"`
	LoginTime util.JSONTime `xorm:"not null default '2006-01-02 15:04:05'" json:"login_time"`
	LoginIp   string        `json:"login_ip"`
}

func NewSAdmin(id int) *SAdmin {
	return &SAdmin{Id: id}
}

// func (p *SAdmin) UnmarshalJSON([]byte) error {
// 	return nil
// }

func (p *SAdmin) TableName() string {
	return "s_admin"
}
