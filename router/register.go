package router

import (
	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {
	registerIndex(r)

	// admin := r.Group("admin")
	// web.Init(r)
	//
	// api := r.Group("api")
	// v1.Init(api.Group("v1"))
}
