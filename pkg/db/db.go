package db

import (
	"errors"

	"github.com/go-xorm/xorm"
)

var (
	defaultDB *xorm.Engine
	dbs       []*xorm.Engine
)

func GetDefaultDB() *xorm.Engine {
	return defaultDB
}

func SetDefaultDB(engine *xorm.Engine) {
	defaultDB = engine
}

func SetDefaultDBByIndex(index uint) {
	defaultDB = GetDBByIndex(index)
}

func SetDBS(options []*Option) error {
	if len(options) == 0 {
		return errors.New("engine options is null")
	}

	for idx := range options {
		engine, err := options[idx].NewEngine()
		if err != nil {
			return err
		}

		dbs = append(dbs, engine)
	}

	defaultDB = dbs[0]

	return nil
}

func GetDBByIndex(index uint) *xorm.Engine {
	nums := uint(len(dbs))
	if nums == 0 {
		return nil
	}
	if nums < (index + 1) {
		return nil
	}

	return dbs[index]
}
