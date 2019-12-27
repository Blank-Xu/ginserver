package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "ginserver/docs"

	defaultInit "ginserver/init"
	"ginserver/routers"
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

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	pid := os.Getpid()

	defer func() {
		if err := recover(); err != nil {
			log.Printf("server pid[%d] crashed with error: %v", pid, err)
			// 等待日志记录完成
			time.Sleep(time.Second)
			panic(err)
		}
		time.Sleep(time.Second)
	}()

	defaultInit.Init()
	routers.Register()

	server := defaultInit.GetConfig().HttpServer.NewServer(nil)
	// run server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server pid[%d] exit with err: %v", pid, err)
		}
	}()
	log.Printf("server pid[%d] start success.", pid)

	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)

	log.Printf("server pid[%d] receive shutdown signal: [%v]", pid, <-quitSignal)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server pid[%d] shutdown failed, err: %v", pid, err)
	}

	log.Printf("server pid[%d] stoped", pid)
}
