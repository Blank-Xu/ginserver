package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/modules/config"
	"ginserver/modules/resp"
	v1 "ginserver/routers/api/v1"
)

func Init(r *gin.Engine) {
	apiRouter := r.Group("api")
	apiRouter.GET("/", info)
	// auth
	apiRouter.Use()

	v1.Init(apiRouter)
}

func info(c *gin.Context) {
	c.JSON(http.StatusOK,
		&resp.Response{
			Code: http.StatusOK,
			Msg:  "api version: " + config.Version,
		})
}
