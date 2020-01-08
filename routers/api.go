package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"ginserver/controllers/v1/system"
)

func registerApi(router *gin.Engine) {
	groupApi := router.Group("api")

	if gin.Mode() != gin.ReleaseMode {
		// register swagger doc router
		groupApi.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		groupApi.GET("/", func(ctx *gin.Context) {
			ctx.Redirect(http.StatusFound, "/swagger/index.html")
		})
	}

	apiV1 := groupApi.Group("v1")
	{
		// apiV1.Use(middlewares.JwtAuth())

		// apiV1.Use(middlewares.CasbinEnforce(casbin.GetEnforcer()))

		user := new(system.UserController)
		apiV1.GET("user/:id", user.GetOne)
		apiV1.GET("user", user.Get)
		apiV1.POST("user", user.Post)
		apiV1.PUT("user/:id", user.Put)
		apiV1.DELETE("user/:id", user.Delete)
	}
}
