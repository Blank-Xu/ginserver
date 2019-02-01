package api

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"

	v1 "ginserver/controllers/api/v1"
	"ginserver/modules/config"
)

func Init(r *gin.Engine, enforcer *casbin.Enforcer) {
	apiRouter := r.Group("api")
	apiRouter.GET("/", info)
	// jwt and casbin auth
	apiRouter.Use(authJwt(enforcer))

	v1.Init(apiRouter)
}

func info(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"version": config.Version})
}
