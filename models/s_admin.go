package models

import (
	"ginserver/modules/db"
)

type SAdmin struct {
	*db.Model  `xorm:"-" json:"-"`
	Id         int         `xorm:"pk autoincr" json:"id"`
	Username   string      `xorm:"unique" json:"username"`
	Password   string      `json:"-"`
	Salt       string      `json:"-"`
	Nickname   string      `json:"nickname"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	Remark     string      `json:"remark"`
	RegisterIp string      `json:"-"`
	Created    db.JSONTime `json:"created" swaggertype:"primitive,integer"`
	Updated    db.JSONTime `json:"updated" swaggertype:"primitive,string"`
	LoginTime  db.JSONTime `json:"login_time" swaggertype:"primitive,string"`
	LoginIp    string      `json:"login_ip"`
}

func NewSAdmin(id int) *SAdmin {
	return &SAdmin{Id: id}
}

func (p *SAdmin) TableName() string {
	return "s_admin"
}

type SAdminInsert struct {
	*db.Model  `xorm:"-" json:"-"`
	Username   string      `xorm:"unique" json:"username" binding:"required"`
	Password   string      `json:"password" binding:"required"`
	Salt       string      `json:"-"`
	Enable     bool        `json:"enable"`
	Nickname   string      `json:"nickname"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	Remark     string      `json:"remark"`
	RegisterIp string      `json:"-"`
	Created    db.JSONTime `xorm:"created" json:"-"`
	Updated    db.JSONTime `xorm:"updated" json:"-"`
}

func (p *SAdminInsert) TableName() string {
	return "s_admin"
}

type SAdminUpdate struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int         `xorm:"pk autoincr" json:"-"`
	Password  string      `json:"password"`
	Salt      string      `json:"-"`
	Enable    bool        `json:"enable"`
	Nickname  string      `json:"nickname"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Remark    string      `json:"remark"`
	Updated   db.JSONTime `xorm:"updated" json:"-"`
}

func (p *SAdminUpdate) TableName() string {
	return "s_admin"
}
