package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/modules/config"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"version": config.Version})
}
