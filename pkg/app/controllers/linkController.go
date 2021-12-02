package controllers

import (
	"shorten-link/pkg/app/models"
	"shorten-link/pkg/app/logic"
	"shorten-link/pkg/db"

	"github.com/go-redis/redis"

	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
)

type Link struct {
	Initial string
}

/*
Receives full link, validates and processes it 
*/
func ShortenHandler(w http.ResponseWriter, r *http.Request, pgClient *db.DB, rdClients []*redis.Client) {
	var receivedLink Link

	err := json.NewDecoder(r.Body).Decode(&receivedLink)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if len(receivedLink.Initial) == 0 {
		http.Error(w, "Can't create shortened version for void", 400)
		return
	}

	validUrl, err := url.ParseRequestURI(receivedLink.Initial)
	if len(validUrl.Host) == 0 || validUrl == nil || err != nil{
		http.Error(w, "Incorrect URL format", 400)
		return
	} 
	
	if pgClient != nil {
		fmt.Println("1")
		existingLinks := models.GetByFullLink(receivedLink.Initial, pgClient)
		fmt.Println("2")
		if len(existingLinks) > 0 {
			fmt.Println("3")
			fmt.Println(existingLinks[0].Hash)
			fmt.Fprintf(w, existingLinks[0].Hash)
		} else {
			fmt.Println("4")
			createdLink := logic.ShortenLink(receivedLink.Initial, pgClient, []*redis.Client{nil, nil})
			fmt.Fprintf(w, createdLink.Hash)
		}
	} else if rdClients[0] != nil && rdClients[1] != nil {
		fmt.Println("rdClient // linkController")
		existingLinks := models.RedisGetByFullLink(receivedLink.Initial, rdClients[0])
		fmt.Println("existingLinks := ", existingLinks)
		if len(existingLinks) > 0 {
			fmt.Fprintf(w, existingLinks)
		} else {
			fmt.Println("createdLink // linkController")
			createdLink := logic.ShortenLink(receivedLink.Initial, nil, rdClients)
			fmt.Fprintf(w, createdLink.Hash)
		}
	}
}

/*
Receives shortened link, validates and redirects to original one
*/
func RedirectLink(w http.ResponseWriter, r *http.Request, pgClient *db.DB, rdClients []*redis.Client) {
	shortenedLink := r.URL.Path[1:]

	if pgClient != nil {
		fullLink := models.GetByShortenedLink(shortenedLink, pgClient)
		if len(fullLink) == 0 {
			http.Error(w, "This link is non-existing or expired", 400)
		} else {
			http.Redirect(w, r, fullLink[0].OriginalUrl, 301)
		}
	} else if rdClients[0] != nil && rdClients[1] != nil {
		fullLink := models.RedisGetByShortenedLink(shortenedLink, rdClients[1])
		fmt.Println("fullLink: ", fullLink)
		if len(fullLink) == 0 {
			http.Error(w, "This link is non-existing or expired", 400)
		} else {
			http.Redirect(w, r, fullLink, 301)
		}
	}
}