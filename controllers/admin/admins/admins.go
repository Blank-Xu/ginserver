package admins

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/e"
)

func GetAdminById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	e.RespErrParamsInvalid(c, (id < 1))
	cols, _ := c.GetQueryArray("cols")

	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record, cols...)
	e.RespErrDBError(c, err)
	e.RespErrNotFound(c, !has)

	e.RespSuccData(c, record)
}
