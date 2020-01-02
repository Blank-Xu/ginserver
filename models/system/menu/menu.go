package s_menu

import (
	"ginserver/pkg/db"
	"ginserver/tools/timeutil"
)

type SMenu struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int `xorm:"pk autoincr"`
	Type      int
	Name      string `xorm:"unique"`
	Method    string
	Path      string
	Icon      string
	Level     int
	OrderNo   int
	State     bool
	ParentId  int
	Created   timeutil.JSONTime `xorm:"created"`
	Updater   int
	Updated   timeutil.JSONTime `xorm:"updated"`
}

func (p *SMenu) TableName() string {
	return "s_menu"
}

func (p *SMenu) SelectByParentId(parentId int) ([]*SMenu, error) {
	var records []*SMenu
	return records, db.GetDefaultDB().SQL(`SELECT *
FROM s_menu
WHERE parent_id = ?
  AND type = ?
  AND state = 1`, parentId, TypeMain).Find(&records)
}
