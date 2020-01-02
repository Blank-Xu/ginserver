package user

import (
	"ginserver/pkg/db"
	"ginserver/tools/timeutil"
)

type User struct {
	*db.Model  `xorm:"-" json:"-"`
	Id         int               `xorm:"pk autoincr" json:"id"`
	Username   string            `xorm:"unique" json:"username"`
	Password   string            `json:"-"`
	Salt       string            `json:"-"`
	State      bool              `json:"state"`
	Nickname   string            `json:"nickname"`
	Icon       string            `json:"icon"`
	Email      string            `json:"email"`
	Phone      string            `json:"phone"`
	Remark     string            `json:"remark"`
	Created    timeutil.JSONTime `json:"created" swaggertype:"primitive,string"`
	Updater    int               `json:"updater"`
	Updated    timeutil.JSONTime `json:"updated" swaggertype:"primitive,string"`
	Deleted    timeutil.JSONTime `json:"-"`
	RegisterIp string            `json:"-"`
	LoginTime  timeutil.JSONTime `json:"login_time" swaggertype:"primitive,string"`
	LoginIp    string            `json:"login_ip"`
}

func (p *User) TableName() string {
	return "s_user"
}

type Insert struct {
	*db.Model  `xorm:"-" json:"-"`
	Username   string            `xorm:"unique" json:"username" binding:"required"`
	Password   string            `json:"password" binding:"required"`
	Salt       string            `json:"-"`
	State      bool              `json:"state"`
	Nickname   string            `json:"nickname"`
	Email      string            `json:"email"`
	Phone      string            `json:"phone"`
	Remark     string            `json:"remark"`
	RegisterIp string            `json:"-"`
	Created    timeutil.JSONTime `xorm:"created" json:"-"`
	Updater    int
	Updated    timeutil.JSONTime `xorm:"updated" json:"-"`
}

func (p *Insert) TableName() string {
	return "s_user"
}

type Update struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int               `xorm:"pk autoincr" json:"-"`
	Password  string            `json:"-"`
	Salt      string            `json:"-"`
	State     bool              `json:"state"`
	Nickname  string            `json:"nickname"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Remark    string            `json:"remark"`
	Updater   int               `json:"-"`
	Updated   timeutil.JSONTime `xorm:"updated" json:"-"`
}

func (p *Update) TableName() string {
	return "s_user"
}

// Login login check
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ChangePwd
type ChangePwd struct {
	Password        string `form:"password" json:"password" binding:"required"`
	NewPassword     string `form:"new_password" json:"new_password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

type InfoUpdate struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int    `xorm:"pk autoincr" json:"-"`
	Nickname  string `form:"nickname" json:"nickname"`
	Icon      string `form:"icon" json:"icon"`
	Email     string `form:"email" json:"email"`
	Phone     string `form:"phone" json:"phone"`
	Updater   int
	Updated   timeutil.JSONTime `xorm:"updated" json:"-"`
}

func (p *InfoUpdate) TableName() string {
	return "s_user"
}
