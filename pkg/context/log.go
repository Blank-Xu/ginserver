package context

import (
	"encoding/json"

	"github.com/rs/zerolog"

	"ginserver/models/log"
)

var (
	logWithoutParamsPath []string
)

func Set(paths []string) {
	logWithoutParamsPath = paths
}

func (p *Context) LogDB(lType log.Type, level zerolog.Level, remark ...string) {
	var (
		params  string
		lRemark string
	)
	for _, v := range logWithoutParamsPath {
		if v == p.Request.URL.Path {
			params = "{}"
			break
		}
	}
	if len(params) == 0 {
		if p.Request.Form != nil {
			param, _ := json.Marshal(p.Request.Form)
			params = string(param)
		} else {
			params = "{}"
		}
	}

	if len(remark) > 0 {
		lRemark = remark[0]
	}

	err := log.Insert(level, lType, p.UserId, p.RoleId, p.Request.Method, p.Request.URL.Path,
		params, p.ClientIP(), lRemark)
	if err != nil {
		p.Error(err)
	}
}
