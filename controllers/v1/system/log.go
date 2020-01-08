package system

import (
	"ginserver/pkg/context"

	"github.com/gin-gonic/gin"
)

type LogController struct {
}

func (*LogController) Get(c *gin.Context) {
	ctx := context.New(c)

	var (
		record  = &logdb.Log{UserId: 0}
		records []*logdb.Log
	)
	if err := record.SelectAll(record, &records); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	// p.Render("log.tpl", map[string]interface{}{"records": records})
}

func (*LogController) Post(c *gin.Context) {
	// ctx := context.New(c)
	//
	// var (
	// 	record  = log.Log{UserId:0}
	// 	records []*log.Log
	// 	err     error
	// )
	// var cond = builder.NewCond()
	// if err = record.Select(&record, &records, cond); err != nil {
	// 	ctx.AbortResponseDatabaseErr(err)
	// 	return
	// }
	//
	// ctx.ResponsePagingOK()
}
