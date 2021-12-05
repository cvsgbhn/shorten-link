package app

import (
	"net/http"
	"os"
	"log"

	"shorten-link/pkg/app/controllers"
	"github.com/go-redis/redis"
	"shorten-link/pkg/db"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func SetupRoutes(dbType string) {

	var pgClient *db.DB
	var rdClients []*redis.Client
	
	switch dbType {
		case "r":
			log.Println("redis")
			rdClients = db.RedisTwoTables()
			log.Println("Connection to db happened:", rdClients)
		case "p":
			log.Println("postgres")
			pgClient = db.NewDB(db.BuildConfig())
			log.Println("Connection to db happened:", pgClient)
		default:
			log.Println("Please, specify database type: redis (r) or postgres (p)")
			return
	}

	port := getPort()

	defer pgClient.Close()

    http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) { 
		controllers.ShortenHandler(w, r, pgClient, rdClients) 
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.RedirectLink(w, r, pgClient, rdClients)
	})

    http.ListenAndServe(port, nil)
}