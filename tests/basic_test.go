package tests

import (
	"shorten-link/pkg/app/controllers"
	"shorten-link/pkg/db"

	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

const (
	DefaultLocalhost = "http://localhost:18001"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

/* Test Cases */
func TestLinkShortenerCRUD(t *testing.T) {
	pgClient := db.NewDB(db.BuildConfig())
	fmt.Println(pgClient)
	getHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.RedirectLink(w, r, pgClient, nil)
	}
	postHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.ShortenHandler(w, r, pgClient, nil)
	}
	t.Run("GET non-existing link", func(t *testing.T) {
		w := performRequest(http.HandlerFunc(getHandler), "GET", DefaultLocalhost + "/shorter/asdfasdfas")

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"This link is non-existing or expired\"}", w.Body.String())
	})
	t.Run("POST empty body", func(t *testing.T) {
		w := performRequest(http.HandlerFunc(postHandler), "POST", DefaultLocalhost + "/shorten/")

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"Incorrect URL format\"}", w.Body.String())
	})
}