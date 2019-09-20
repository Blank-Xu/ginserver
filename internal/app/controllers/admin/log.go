package admin

import (
	"ginserver/internal/app/models/log"

	"github.com/gin-gonic/gin"
)

type ControllerLog struct {
	Controller
}

func (p *ControllerLog) Get(ctx *gin.Context) {
	p.New(ctx)
	var (
		record  = &log.Log{UserId: p.GetUserId()}
		records []*log.Log
	)
	if err := record.SelectAll(record, &records); err != nil {
		p.RespErrDBError(err)
		return
	}
	p.Render("log.tpl", map[string]interface{}{"records": records})
}

func (p *ControllerLog) Post(ctx *gin.Context) {
	// p.New(ctx)
	// var (
	// 	record  = log.Log{UserId: p.GetUserId()}
	// 	records []*log.Log
	// 	err     error
	// )
	// var cond = builder.NewCond()
	// if err = record.SelectCond(&record, &records, cond); err != nil {
	// 	p.RespErrDBError(err)
	// 	return
	// }
	// p.RespOk(records)
}
