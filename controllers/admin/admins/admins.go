package admins

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/e"
)

type Admins struct{}

func (p *Admins) RegisterRouter(r *gin.RouterGroup) {
	r.GET("admins/:id", p.GetAdminById)
}

func (p *Admins) GetAdminById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}
	cols, _ := c.GetQueryArray("cols")

	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record, cols...)
	if err != nil {
		e.RespErrDBError(c, err)
		return
	}
	if !has {
		e.RespErrNotFound(c)
		return
	}

	e.RespSuccData(c, record)
}
