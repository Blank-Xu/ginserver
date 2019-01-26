package db

import (
	"fmt"
	"time"

	"ginserver/modules/config"
	"ginserver/modules/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var defaultEngine *xorm.Engine

func Init() {
	var cfgs = config.GetConfig().DataBase
	defaultEngine = newEngine(cfgs[0])
}

func newEngine(cfg *config.DataBase) *xorm.Engine {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase, cfg.Charset)
	engine, err := xorm.NewEngine(cfg.DriverName, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("database [%s] engine create error \n - dns: [%s] \n - err: [%v]", cfg.DataBase, dataSourceName, err))
	}

	engine.SetLogger(newSimpleLogger(cfg.DataBase, core.LogLevel(cfg.LogLevel)))

	// check connection
	if cfg.Connect {
		if err = engine.Ping(); err != nil {
			panic(fmt.Sprintf("database [%s] connect error \n - dns: [%s] \n - err: [%v]", cfg.DataBase, dataSourceName, err))
		}
		log.Infof("database [%s] connected.", cfg.DataBase)
	}

	engine.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	engine.SetMaxIdleConns(cfg.MaxIdleConns)
	engine.SetMaxOpenConns(cfg.MaxOpenConns)
	engine.ShowExecTime(cfg.ShowExecTime)
	engine.ShowSQL(cfg.ShowSql)
	return engine
}

func GetDefaultEngine() *xorm.Engine {
	return defaultEngine
}
