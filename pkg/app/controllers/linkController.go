package controllers

import (
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

	fmt.Fprintf(w, receivedLink.Initial)
}