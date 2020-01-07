package param

import (
	"ginserver/pkg/db"
	"time"
)

type Param struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int       `xorm:"pk autoincr" json:"id"`
	Name      string    `xorm:"unique"`
	Param     string    `json:"param"`
	Remark    string    `json:"remark"`
	Updated   time.Time `json:"updated"`
}

func (p *Param) TableName() string {
	return "s_param"
}
