package models

import (
	"time"

	"ginserver/tools/db"
)

type SUser struct {
	*db.Model  `xorm:"-" json:"-"`
	Id         int         `xorm:"pk autoincr" json:"id"`
	Username   string      `xorm:"unique" json:"username"`
	Password   string      `json:"-"`
	Salt       string      `json:"-"`
	State      bool        `json:"state"`
	Nickname   string      `json:"nickname"`
	Icon       string      `json:"icon"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	Remark     string      `json:"remark"`
	Created    db.JSONTime `json:"created" swaggertype:"primitive,string"`
	Updater    int         `json:"updater"`
	Updated    db.JSONTime `json:"updated" swaggertype:"primitive,string"`
	Deleted    time.Time   `json:"-"`
	RegisterIp string      `json:"-"`
	LoginTime  db.JSONTime `json:"login_time" swaggertype:"primitive,string"`
	LoginIp    string      `json:"login_ip"`
}

func NewSUser(id int) *SUser {
	return &SUser{Id: id}
}

func (p *SUser) TableName() string {
	return "s_user"
}

type SUserInsert struct {
	*db.Model  `xorm:"-" json:"-"`
	Username   string      `xorm:"unique" json:"username" binding:"required"`
	Password   string      `json:"password" binding:"required"`
	Salt       string      `json:"-"`
	State      bool        `json:"state"`
	Nickname   string      `json:"nickname"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	Remark     string      `json:"remark"`
	RegisterIp string      `json:"-"`
	Created    db.JSONTime `xorm:"created" json:"-"`
	Updater    int
	Updated    db.JSONTime `xorm:"updated" json:"-"`
}

func (p *SUserInsert) TableName() string {
	return "s_user"
}

type SUserUpdate struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int    `xorm:"pk autoincr" json:"-"`
	Password  string `json:"-"`
	Salt      string `json:"-"`
	State     bool   `json:"state"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Remark    string `json:"remark"`
	Updater   int
	Updated   db.JSONTime `xorm:"updated" json:"-"`
}

func (p *SUserUpdate) TableName() string {
	return "s_user"
}

// SUserLogin login check
type SUserLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// SUserChangePwd
type SUserChangePwd struct {
	Password        string `form:"password" json:"password" binding:"required"`
	NewPassword     string `form:"new_password" json:"new_password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

type SUserInfoUpdate struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int    `xorm:"pk autoincr" json:"-"`
	Nickname  string `json:"nickname"`
	Icon      string `json:"icon"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Remark    string `json:"remark"`
	Updater   int
	Updated   db.JSONTime `xorm:"updated" json:"-"`
}
