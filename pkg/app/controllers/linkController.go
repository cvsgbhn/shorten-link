package controllers

import (
	"shorten-link/pkg/app/models"

	"encoding/json"
	"net/http"
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

	existingLinks := models.GetFullLink(receivedLink.Initial)

	if len(existingLinks) > 0 {
		fmt.Fprintf(w, existingLinks[0].Hash)
	} else {
		//createdLink := logic.ShortenLink()
		//fmt.Fprintf(w, createdLink.Hash)
	}

}