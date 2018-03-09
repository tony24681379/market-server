package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/golang/glog"
	"github.com/tony24681379/market-server/config"
	"github.com/tony24681379/market-server/router"
	"github.com/tony24681379/market-server/rtMart"
	"github.com/tony24681379/market-server/shopping"
)

func main() {
	configs := config.NewConfig()
	rtMart := rtMart.NewRtMart()
	shopping := shopping.NewShopping()

	router := router.NewRouter(rtMart, shopping)

	glog.Info("serve port", configs.Port)
	server := &http.Server{
		Addr:    configs.Port,
		Handler: router,
	}
	gracefulShutdown(server)
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			glog.Info("Server closed under request")
		} else {
			glog.Fatal("Server closed unexpect")
		}
	}

	glog.Info("Server exiting")
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()
}
