package app

import (
	"net/http"
	"os"
	"log"

	"shorten-link/pkg/app/controllers"
)

/* get port from environment */
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

	/* Upload CSV action */
    http.HandleFunc("/shorten", controllers.ShortenLink)

    http.ListenAndServe(port, nil)
}