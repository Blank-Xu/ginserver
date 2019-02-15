package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "ginserver/docs"
	defaultInit "ginserver/init"
	"ginserver/init/config"
	"ginserver/internal/app/routers"

	"github.com/sirupsen/logrus"
)

// @title ginserver Swagger API
// @version 0.1
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://github.com/Blank-Xu/ginserver/blob/master/LICENSE

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic JWT Auth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	defaultInit.Init()
	routers.Register()

	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("Server Crashed With err: [%v]", err)
			logrus.Error(msg)
			panic(msg)
		}
	}()

	var cfg = config.GetConfig().Server
	srv := &http.Server{
		Addr:           ":" + cfg.HttpPort,
		Handler:        routers.GetDefaultRouter(),
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("server listen err: ", err)
		}
	}()
	logrus.Info("server start success.")

	quite := make(chan os.Signal)
	signal.Notify(quite, os.Interrupt)
	logrus.Infof("server shutdown with signal %v", <-quite)
	// wait for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// graceful close, need golang version 1.8+
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("shutdown err: ", err)
	}
	logrus.Info("server exit.")
}
