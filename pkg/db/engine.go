package db

import (
	"github.com/go-xorm/xorm"
)

var (
	defaultEngine *xorm.Engine
	engineSlice   *[]*xorm.Engine
)

func GetDefaultEngine() *xorm.Engine {
	return defaultEngine
}

func SetDefaultEngine(engine *xorm.Engine) {
	defaultEngine = engine
}

func SetDefaultEngineByIndex(index uint) {
	defaultEngine = GetEngineByIndex(index)
}

func SetEngines(engines *[]*xorm.Engine) {
	engineSlice = engines
	if len(*engines) > 0 {
		defaultEngine = (*engines)[0]
	}
}

func GetEngineByIndex(index uint) *xorm.Engine {
	nums := uint(len(*engineSlice))
	if nums == 0 {
		return nil
	}
	if nums < (index + 1) {
		return nil
	}
	return (*engineSlice)[index]
}
