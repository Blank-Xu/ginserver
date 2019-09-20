package init

import (
	"fmt"
	"time"

	"ginserver/init/config"
	"ginserver/tools/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

func dbInit() {
	var cfgDatabase = config.GetConfig().DataBase

	engines := make([]*xorm.Engine, len(cfgDatabase))
	for idx := range cfgDatabase {
		engines[idx] = newEngine(cfgDatabase[idx])
	}
	db.SetEngines(&engines)
	// db.SetEngines will set default engine by index 0
	// db.SetDefaultEngineByIndex(0)
}

func newEngine(cfg *config.DataBase) *xorm.Engine {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase, cfg.Charset)

	engine, err := xorm.NewEngine(cfg.DriverName, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("Database [%s] engine create error \n - dns: [%s] \n - err: [%v]", cfg.DataBase, dataSourceName, err))
	}

	engine.SetLogger(db.NewSimpleLogger(logrus.StandardLogger(), cfg.DataBase, core.LogLevel(cfg.LogLevel)))

	// check connection
	if cfg.Connect {
		if err = engine.Ping(); err != nil {
			panic(fmt.Sprintf("Database [%s] connect error \n - dns: [%s] \n - err: [%v]", cfg.DataBase, dataSourceName, err))
		}
		logrus.Infof("database [%s] connected.", cfg.DataBase)
	}

	engine.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	engine.SetMaxIdleConns(cfg.MaxIdleConns)
	engine.SetMaxOpenConns(cfg.MaxOpenConns)
	engine.ShowExecTime(cfg.ShowExecTime)
	engine.ShowSQL(cfg.ShowSql)
	return engine
}
