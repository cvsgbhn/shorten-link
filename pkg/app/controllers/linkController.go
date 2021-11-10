package controllers

import (
	"shorten-link/pkg/app/models"
	"shorten-link/pkg/app/logic"

	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
)

type Link struct {
	Initial string
}

func ShortenLink(w http.ResponseWriter, r *http.Request) {
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
	
	existingLinks := models.GetByFullLink(receivedLink.Initial)

	if len(existingLinks) > 0 {
		fmt.Fprintf(w, existingLinks[0].Hash)
	} else {
		createdLink := logic.ShortenLink(receivedLink.Initial)
		fmt.Fprintf(w, createdLink.Hash)
	}
}

func RedirectLink(w http.ResponseWriter, r *http.Request) {
	shortenedLink := r.URL.Path[1:]

	fullLink := models.GetByShortenedLink(shortenedLink)
	if len(fullLink) == 0 {
		http.Error(w, "This link is non-existing or expired", 400)
	} else {
		http.Redirect(w, r, fullLink[0].OriginalUrl, 301)
	}
}