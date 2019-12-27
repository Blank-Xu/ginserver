package db

import (
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

func SetDBs(engines []*xorm.Engine) {
	dbs = engines
	if len(engines) > 0 {
		defaultDB = engines[0]
	}
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
