package main

import (
	"dreamland/pkg/router"
)

func main() {
	engine := router.Default()
	engine.Run()
}
