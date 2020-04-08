package main

import (
	"dreamland/pkg/db"
	"dreamland/pkg/router"
)

func main() {
	db.InitDB()
	engine := router.Default()
	engine.Run()
}
