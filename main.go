package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ginserver/controllers"
	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/db"
	"ginserver/modules/log"

	"github.com/sirupsen/logrus"
)

var configFile = flag.String("config", "config/app_debug.yaml", "config file")

func init() {
	flag.Parse()
	fmt.Printf("server starting ... \n - version: [%s]  \n - args: %s\n", config.Version, os.Args)
	fmt.Printf("read config file ... \n - file_name: [%s]\n", *configFile)
	fmt.Println(" - you can use [-config file] command to set config file when server start.")

	config.Init(*configFile)
	fmt.Println("read config success")

	// fix default setting
	fix()

	log.Init()
	logrus.Info("server starting ...")

	db.Init()

	models.Init()

	controllers.Init()
}

func main() {
	var cfg = config.GetConfig().Server

	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("server crashed with err: [%v]", err)
			logrus.Error(msg)
			panic(msg)
		}
	}()

	srv := &http.Server{
		Addr:           ":" + cfg.HttpPort,
		Handler:        controllers.GetRouter(),
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

func fix() {
	var cfg = config.GetConfig().Fix
	// fix timezone
	time.Local = time.FixedZone(cfg.TimeZone.Name, cfg.TimeZone.Offset*3600)
}
