package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/multitemplate"

	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/e"
	"ginserver/modules/func_map"
	"ginserver/modules/middleware"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const staticFile = "/assets/"

var router = gin.New()

func Init() {
	var cfg = config.GetConfig()
	// set global setting
	if cfg.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
	}
	gin.SetMode(cfg.RunMode)

	// set global middleware
	router.Use(middleware.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())
	// statics and templates
	router.Static(cfg.StaticFile, staticFile)
	router.SetFuncMap(func_map.GetFunMap())
	router.HTMLRender = loadTemplates(cfg.TemplateFile)
	// cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
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
	casbinEnforcer := casbin.NewEnforcer(cfg.RbacFile, &models.SCasbin{})
	// register routers
	registerRouter(router, casbinEnforcer)
}

func GetRouter() *gin.Engine {
	return router
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Layouts err: [%v]", err))
	}

	includes, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
