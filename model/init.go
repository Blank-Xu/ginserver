package model

import "ginserver/module/db"

func Init() {
	var xorm = db.GetDefaultEngine()
	if err := xorm.Sync2(new(Casbin)); err != nil {
		panic("xorm sync err: " + err.Error())
	}
}
