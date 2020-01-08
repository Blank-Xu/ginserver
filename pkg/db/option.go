package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/rs/zerolog/log"

	"xorm.io/core"
)

type Option struct {
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

func (p *Option) NewEngine() (*xorm.Engine, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		p.Username, p.Password, p.Host, p.Port, p.DataBase, p.Charset)

	engine, err := xorm.NewEngine(p.DriverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("xorm engine create failed, database: [%s], source: [%s], err: %v", p.DataBase, dataSourceName, err)
	}

	engine.SetLogger(NewSimpleLogger(&log.Logger, p.DataBase, core.LogLevel(p.LogLevel)))

	// check connection
	if p.Connect {
		if err = engine.Ping(); err != nil {
			return nil, fmt.Errorf("xorm engine Ping failed, database: [%s], err: %v", p.DataBase, err)
		}
	}

	engine.SetConnMaxLifetime(time.Minute * time.Duration(p.ConnMaxLifetime))
	engine.SetMaxIdleConns(p.MaxIdleConns)
	engine.SetMaxOpenConns(p.MaxOpenConns)
	engine.ShowExecTime(p.ShowExecTime)
	engine.ShowSQL(p.ShowSql)

	return engine, nil
}
