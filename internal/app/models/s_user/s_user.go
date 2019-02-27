package s_user

import (
	"time"

	"ginserver/tools/db"
)

type User struct {
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

func (p *User) TableName() string {
	return "s_user"
}

type UserInsert struct {
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

func (p *UserInsert) TableName() string {
	return "s_user"
}

type UserUpdate struct {
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

func (p *UserUpdate) TableName() string {
	return "s_user"
}

// UserLogin login check
type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserChangePwd
type UserChangePwd struct {
	Password        string `form:"password" json:"password" binding:"required"`
	NewPassword     string `form:"new_password" json:"new_password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

type UserInfoUpdate struct {
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
