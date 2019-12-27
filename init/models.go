package init

import (
	"fmt"

	"github.com/go-xorm/xorm"

	"ginserver/pkg/db"
)

func modelsInit() {
	var engine = db.GetDefaultEngine()

	// just need to synchronize auto create tables
	if err := syncServerTables(engine); err != nil {
		panic(fmt.Sprintf("Database [%s] server tables synchronize err: [%v]", engine.DataSourceName(), err))
	}
}

// syncServerTables synchronize server manage tables
func syncServerTables(engine *xorm.Engine) error {
	return engine.Sync2()
}
