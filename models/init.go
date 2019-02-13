package models

import (
	"fmt"

	"github.com/go-xorm/xorm"

	"ginserver/modules/db"
)

func Init() {
	var engine = db.GetDefaultEngine()

	// just need to synchronize auto create tables
	if err := syncServerTables(engine); err != nil {
		panic(fmt.Sprintf("database [%s] server tables synchronize err: [%v]", engine.DataSourceName(), err))
	}
}

// syncServerTables synchronize server manage tables
func syncServerTables(engine *xorm.Engine) error {
	return engine.Sync2()
}
