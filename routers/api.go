package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"ginserver/controllers/v1/system"
)

func registerApi(router *gin.Engine) {
	// register swagger doc router
	router.GET("swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	groupApi := router.Group("api")
	groupApi.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	apiV1 := groupApi.Group("v1")
	{
		// apiV1.Use(middlewares.JwtAuth())

		// apiV1.Use(middlewares.CasbinEnforce(casbin.GetEnforcer()))

		user := new(system.UserController)
		apiV1.GET("users/:id", user.GetOne)
		apiV1.GET("users", user.Get)
		apiV1.POST("users", user.Post)
		apiV1.PUT("users/:id", user.Put)
		apiV1.DELETE("users/:id", user.Delete)
	}
}
