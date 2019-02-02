package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"ginserver/models"
	"ginserver/modules/e"
)

type admin struct{}

func (p *admin) registerRouter(r *gin.RouterGroup) {
	r.GET("admins/:id", p.get)
}

func (p *admin) get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}
	cols, _ := c.GetQueryArray("cols")
	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record, cols...)
	if err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	if !has {
		e.RespErrNotFound(c)
		return
	}
	e.RespDataOk(c, record)
}
