package models

import (
	"ginserver/modules/db"
	"ginserver/modules/utils"
)

type SAdmin struct {
	*db.Model  `xorm:"-" json:"-"`
	Id         int            `xorm:"pk autoincr" json:"id"`
	Username   string         `xorm:"unique" json:"username"`
	Password   string         `json:"-"`
	Salt       string         `json:"-"`
	Nickname   string         `json:"nickname"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Remark     string         `json:"remark"`
	RegisterIp string         `json:"-"`
	Created    utils.JSONTime `json:"created"`
	Updated    utils.JSONTime `json:"updated"`
	LoginTime  utils.JSONTime `json:"login_time"`
	LoginIp    string         `json:"login_ip"`
}

const (
	SAdminInsertCols = "username,password,salt,nickname,email,phone,remark,register_ip"
)

func NewSAdmin(id int) *SAdmin {
	return &SAdmin{Id: id}
}

func (p *SAdmin) TableName() string {
	return "s_admin"
}

func (p *SAdmin) Select(orderBy string, paging *db.Paging, cols ...string) (records []*SAdmin, err error) {
	err = db.GetDefaultEngine().Cols(cols...).OrderBy(orderBy).Limit(paging.LimitOffset()).Find(&records)
	return
}
