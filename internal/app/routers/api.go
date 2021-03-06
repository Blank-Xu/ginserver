package routers

import (
	"net/http"

	"ginserver/internal/app/controllers/api/v1/users"
	"ginserver/tools/casbin"
	"ginserver/tools/middleware"

	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func registerApiRouter(router *gin.Engine) {
	// register swagger doc router
	router.GET("swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	groupApi := router.Group("api")
	groupApi.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	apiV1 := groupApi.Group("v1")
	{
		apiV1.Use(middleware.JwtAuth())

		apiV1.Use(middleware.CasbinEnforce(casbin.GetEnforcer()))
		apiV1.GET("users/:id", new(users.ControllerUsers).GetOne)
		apiV1.GET("users", new(users.ControllerUsers).Get)
		apiV1.POST("users", new(users.ControllerUsers).Post)
		apiV1.PUT("users/:id", new(users.ControllerUsers).Put)
		apiV1.DELETE("users/:id", new(users.ControllerUsers).Delete)
	}
}
