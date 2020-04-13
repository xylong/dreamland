package main

import (
	"dreamland/config"
	"dreamland/pkg/db"
	"dreamland/pkg/router"
)

func main() {
	config.InitConfig()
	db.InitDB()
	engine := router.Default()
	engine.Run()
}
