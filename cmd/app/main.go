package main

import (
	"shorten-link/pkg/app"

	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Please, specify database type: redis (r) or postgres (p)")
		return
	}
	app.SetupRoutes(os.Args[1])
}
