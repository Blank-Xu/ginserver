package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ginserver/module/config"
	"ginserver/module/db"
	glog "ginserver/module/log"
	"ginserver/router"
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

	glog.Init()
	glog.Info("server starting ...")

	db.Init()

	router.Init()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("server crashed with err: [%v]", err)
			glog.Error(msg)
			panic(msg)
		}
	}()

	var cfg = config.GetConfig().Server
	var defaultLog = glog.GetLog()

	srv := &http.Server{
		Addr:           ":" + cfg.HttpPort,
		Handler:        router.GetRouter(),
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       log.New(defaultLog.Writer(), "http", 0),
	}

	// run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			defaultLog.Fatal("server listen err: ", err)
		}
	}()
	defaultLog.Info("server start success.")

	quite := make(chan os.Signal)
	signal.Notify(quite, os.Interrupt)
	defaultLog.Infof("server shutdown with signal %v", <-quite)

	// wait for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful close, need golang version 1.8+
	if err := srv.Shutdown(ctx); err != nil {
		defaultLog.Fatal("shutdown err: ", err)
	}
	defaultLog.Info("server exit.")
}

func fix() {
	var cfg = config.GetConfig().Fix
	// fix timezone
	time.Local = time.FixedZone(cfg.TimeZone.Name, cfg.TimeZone.Offset*3600)
}
