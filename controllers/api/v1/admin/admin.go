package admin

import (
	"net/http"
	"strconv"

	"github.com/go-xorm/builder"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/errors"
	"ginserver/modules/resp"
)

func GetAdminById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(http.StatusBadRequest, errors.RespError(errors.CodeParamInvalid))
		return
	}

	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record)
	if err != nil {
		c.JSON(http.StatusNotImplemented, errors.RespError(errors.CodeDBErr, err))
		return
	}
	if !has {
		c.JSON(errors.RespHttpError(http.StatusNotFound))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: &record})
}

func GetAdmin(c *gin.Context) {
	var (
		record  = new(models.SAdmin)
		records []*models.SAdmin
	)
	if err := record.Select(record, &records); err != nil {
		c.JSON(http.StatusNotImplemented, errors.RespError(errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: &records})
}

func PostAdmin(c *gin.Context) {
	record := new(models.SAdmin)
	if err := c.ShouldBind(record); err != nil {
		c.JSON(http.StatusBadRequest, errors.RespError(errors.CodeParamInvalid, err))
		return
	}

	record.RegisterIp = c.ClientIP()
	count, err := record.InsertOne(record)
	if err != nil {
		c.JSON(http.StatusNotImplemented, errors.RespError(errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}

func PutAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(http.StatusBadRequest, errors.RespError(errors.CodeParamInvalid))
		return
	}
	record := new(models.SAdmin)
	cond := builder.Gt{"id": id}
	has, err := record.IsRecordExists("id", record, cond)
	if err != nil {
		c.JSON(http.StatusNotImplemented, errors.RespError(errors.CodeDBErr, err))
		return
	}
	if !has {
		c.JSON(errors.RespHttpError(http.StatusNotFound))
		return
	}

	if err := c.ShouldBind(record); err != nil {
		c.JSON(http.StatusBadRequest, errors.RespError(errors.CodeParamInvalid, err))
		return
	}

	count, err := record.Update(record, cond)
	if err != nil {
		c.JSON(http.StatusNotImplemented, errors.RespError(errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}

func DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(http.StatusBadRequest, errors.RespError(errors.CodeParamInvalid))
		return
	}

	record := models.NewSAdmin(id)
	count, err := record.Delete(record)
	if err != nil {
		c.JSON(http.StatusNotImplemented, errors.RespError(errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}
