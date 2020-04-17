package main

import (
	"context"
	"dreamland/config"
	"dreamland/pkg/db"
	"dreamland/pkg/router"
	"dreamland/pkg/util"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitConfig()
	db.InitDB()
	util.InitValidate()
	engine := router.Default()

	serve := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("server.port")),
		Handler: engine,
	}

	go func() {
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := serve.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}
	log.Println("server exiting...")
}
