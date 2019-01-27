package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/errors"
	"ginserver/modules/resp"
)

func GetAdminById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(errors.RespError(http.StatusBadRequest, errors.CodeParamInvalid))
		return
	}

	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record)
	if err != nil {
		c.JSON(errors.RespError(http.StatusNotImplemented, errors.CodeDBErr, err))
		return
	}
	if !has {
		c.JSON(errors.RespError(http.StatusBadRequest, errors.CodeParamInvalid))
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
		c.JSON(errors.RespError(http.StatusNotImplemented, errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: &records})
}

func PostAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(errors.RespError(http.StatusBadRequest, errors.CodeParamInvalid))
		return
	}
	// var record  = new(models.SAdmin)
}

func PutAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(errors.RespError(http.StatusBadRequest, errors.CodeParamInvalid))
		return
	}

	record := models.NewSAdmin(id)
	count, err := record.Update(record)
	if err != nil {
		c.JSON(errors.RespError(http.StatusNotImplemented, errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}

func DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(errors.RespError(http.StatusBadRequest, errors.CodeParamInvalid))
		return
	}

	record := models.NewSAdmin(id)
	count, err := record.Delete(record)
	if err != nil {
		c.JSON(errors.RespError(http.StatusNotImplemented, errors.CodeDBErr, err))
		return
	}
	c.JSON(http.StatusOK, &resp.ResponseData{Data: count})
}
