package models

import (
	"time"

	"ginserver/modules/db"
)

type SAdmin struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int       `xorm:"pk autoincr" json:"id"`
	Username  string    `xorm:"unique" json:"username"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Remark    string    `json:"remark"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	LoginTime time.Time `json:"login_time"`
	LoginIp   string    `json:"login_ip"`
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
