package routers

import (
	"github.com/gin-gonic/gin"

	"ginserver/routers/api"
	"ginserver/routers/web"
)

func registerRouter(r *gin.Engine) {
	registerIndex(r)

	web.Init(r)

	api.Init(r)
}
