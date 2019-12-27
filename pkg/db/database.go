package db

import (
	"fmt"
	"time"
	
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

type Database struct {
	DriverName      string `yaml:"DriverName"`
	DataBase        string `yaml:"DataBase"`
	Host            string `yaml:"Host"`
	Port            string `yaml:"Port"`
	Username        string `yaml:"Username"`
	Password        string `yaml:"Password"`
	Charset         string `yaml:"Charset"`
	LogLevel        int    `yaml:"LogLevel"`
	ConnMaxLifetime int    `yaml:"MaxLifetime"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	ShowSql         bool   `yaml:"ShowSql"`
	ShowExecTime    bool   `yaml:"ShowExecTime"`
	Connect         bool   `yaml:"Connect"`
}

func (p *Database)NewEngine()(*xorm.Engine, error){
	
	return nil,nil
}

func (p *Database) Init() (*xorm.Engine, error) {
	// engines := make([]*xorm.Engine, len(cfgDatabase))
	// for idx := range cfgDatabase {
	// 	engines[idx] = newEngine(cfgDatabase[idx])
	// }
	// db.SetEngines(&engines)
	// db.SetEngines will set default engine by index 0
	// db.SetDefaultDBByIndex(0)
	
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		p.Username, p.Password, p.Host, p.Port, p.DataBase, p.Charset)
	
	engine, err := xorm.NewEngine(p.DriverName, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("Database [%s] engine create error \n - dns: [%s] \n - err: [%v]", p.DataBase, dataSourceName, err))
	}
	
	engine.SetLogger(NewSimpleLogger(logrus.StandardLogger(), p.DataBase, core.LogLevel(p.LogLevel)))
	
	// check connection
	if p.Connect {
		if err = engine.Ping(); err != nil {
			panic(fmt.Sprintf("Database [%s] connect error \n - dns: [%s] \n - err: [%v]", p.DataBase, dataSourceName, err))
		}
		logrus.Infof("database [%s] connected.", p.DataBase)
	}
	
	engine.SetConnMaxLifetime(time.Duration(p.ConnMaxLifetime) * time.Minute)
	engine.SetMaxIdleConns(p.MaxIdleConns)
	engine.SetMaxOpenConns(p.MaxOpenConns)
	engine.ShowExecTime(p.ShowExecTime)
	engine.ShowSQL(p.ShowSql)
	return engine, nil
}