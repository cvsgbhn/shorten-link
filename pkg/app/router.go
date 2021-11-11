package app

import (
	"net/http"
	"os"
	"log"

	"shorten-link/pkg/app/controllers"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func SetupRoutes() {
	port := getPort()

    http.HandleFunc("/shorten", controllers.ShortenHandler)
	http.HandleFunc("/", controllers.RedirectLink)

    http.ListenAndServe(port, nil)
}