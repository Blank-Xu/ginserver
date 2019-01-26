package routers

import (
	"github.com/gin-gonic/gin"

	"ginserver/controllers"
	"ginserver/routers/admin"
	"ginserver/routers/api"
)

func registerRouter(r *gin.Engine) {
	controllers.GetIndex(r)

	admin.Init(r)

	api.Init(r)
}
