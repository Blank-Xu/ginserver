package param

import (
	"ginserver/pkg/db"
	"ginserver/tools/timeutil"
)

type Param struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int               `xorm:"pk autoincr" json:"id"`
	Name      string            `xorm:"unique"`
	Param     string            `json:"param"`
	Remark    string            `json:"remark"`
	Updated   timeutil.JSONTime `json:"updated"`
}

func (p *Param) TableName() string {
	return "s_param"
}
