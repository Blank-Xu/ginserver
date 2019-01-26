package model

import (
	"fmt"

	"ginserver/module/db"
)

func Init() {
	var orm = db.GetDefaultEngine()
	if err := orm.Sync2(
		new(SCasbin),
	); err != nil {
		panic(fmt.Sprintf("database [%s] synchronize err: [%v]", orm.DataSourceName(), err))
	}
}
