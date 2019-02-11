package routers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"ginserver/controllers"
	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/e"
	"ginserver/modules/func_map"
	"ginserver/modules/middleware"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	router   = gin.New()
	enforcer *casbin.Enforcer
)

func Init() {
	var cfg = config.GetConfig()
	// set global setting
	gin.SetMode(cfg.RunMode)
	if cfg.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
		gin.DefaultErrorWriter = logrus.StandardLogger().Out
		router.Use(middleware.Logger(logrus.StandardLogger()))
	} else {
		router.Use(gin.Logger())
	}

	// set global middleware
	router.Use(gin.Recovery())
	// statics and templates
	router.Static("/statics", cfg.StaticDir)
	router.SetFuncMap(func_map.GetFunMap())
	router.HTMLRender = loadTemplates(cfg.TemplateDir)
	// cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusMethodNotAllowed))
	})

	// load casbin
	enforcer = casbin.NewEnforcer(cfg.RbacFile, &models.SCasbin{})
	// register swagger doc router
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// home page
	router.GET("/", new(controllers.IndexController).Get)

	registerAdminRouter()

	registerApiRouter()
}

func GetRouter() *gin.Engine {
	return router
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Layouts err: [%v]", err))
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tpl")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}

	includes2, err := filepath.Glob(templatesDir + "/includes/**/*.tpl")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}
	includes = append(includes, includes2...)
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	templates, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(fmt.Sprintf("Load Templates err: [%v]", err))
	}
	for _, template := range templates {
		r.AddFromFiles(filepath.Base(template), template)
	}
	return r
}
