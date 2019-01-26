package models

import (
	"fmt"

	"github.com/go-xorm/xorm"

	"ginserver/modules/db"
)

func Init() {
	var orm = db.GetDefaultEngine()

	// just need to synchronize auto create tables
	if err := syncServerTables(orm); err != nil {
		panic(fmt.Sprintf("database [%s] server tables synchronize err: [%v]", orm.DataSourceName(), err))
	}

	if err := syncBusinessTables(orm); err != nil {
		panic(fmt.Sprintf("database [%s] business tables synchronize err: [%v]", orm.DataSourceName(), err))
	}
}

// syncServerTables synchronize server manage tables
func syncServerTables(orm *xorm.Engine) error {
	return orm.Sync2(
		new(SCasbin),
	)
}

// syncBusinessTables synchronize business tables
func syncBusinessTables(orm *xorm.Engine) error {
	return orm.Sync2()
}
