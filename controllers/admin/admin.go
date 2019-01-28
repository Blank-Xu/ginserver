package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/e"
)

func GetAdminById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(http.StatusBadRequest, e.RespError(e.CodeParamInvalid))
		return
	}

	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record)
	if err != nil {
		c.JSON(http.StatusNotImplemented, e.RespError(e.CodeDBErr, err))
		return
	}
	if !has {
		c.JSON(http.StatusBadRequest, e.RespError(e.CodeParamInvalid))
		return
	}
	c.JSON(http.StatusOK, record)
}

func GetAdmin(c *gin.Context) {

}

func PostAdmin(c *gin.Context) {

}

func PutAdmin(c *gin.Context) {

}

func DeleteAdmin(c *gin.Context) {

}
