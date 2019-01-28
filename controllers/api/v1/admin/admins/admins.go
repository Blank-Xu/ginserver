package admins

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/builder"

	"ginserver/models"
	"ginserver/modules/e"
	"ginserver/modules/log"
	"ginserver/modules/resp"
)

func GetAdminById(c *gin.Context) {
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
		log.Error(err)
		return
	}
	if !has {
		e.RespErrNotFound(c)
		return
	}
	e.RespSuccData(c, record)
}

func GetAdmin(c *gin.Context) {
	var (
		record  = new(models.SAdmin)
		records []*models.SAdmin
	)
	err := record.Select(record, &records)
	if err != nil {
		e.RespErrDBError(c, err)
		log.Error(err)
		return
	}

	e.RespSuccData(c, &records)
}

func PostAdmin(c *gin.Context) {
	record := new(models.SAdmin)
	if err := c.ShouldBind(record); err != nil {
		c.JSON(http.StatusBadRequest, e.RespErrCode(e.CodeParamInvalid, err))
		return
	}

	record.RegisterIp = c.ClientIP()
	count, err := record.InsertOne(record)
	if err != nil {
		c.JSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}

func PutAdmin(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}
	record := new(models.SAdmin)
	cond := builder.Gt{"id": id}
	has, err := record.IsRecordExists("id", record, cond)
	if err != nil {
		c.JSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr, err))
		return
	}
	if !has {
		c.JSON(e.RespErrHttp(http.StatusNotFound))
		return
	}

	if err := c.ShouldBind(record); err != nil {
		c.JSON(http.StatusBadRequest, e.RespErrCode(e.CodeParamInvalid, err))
		return
	}

	count, err := record.Update(record, cond)
	if err != nil {
		c.JSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}

func DeleteAdmin(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}

	record := models.NewSAdmin(id)
	count, err := record.Delete(record)
	if err != nil {
		e.RespErrDBError(c, err)
		log.Error(err)
		return
	}

	e.RespSuccData(c, count)
}
