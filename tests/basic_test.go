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
	DefaultLocalhost = "http://localhost:18003"
)

/* Test Cases */
func TestLinkShortenerCRUD(t *testing.T) {
	pgClient := db.NewDB(db.BuildConfig())
	fmt.Println(pgClient)
	getHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.RedirectLink(w, r, pgClient, nil)
	}
	// postHandler := func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.ShortenHandler(w, r, pgClient, nil)
	// }
	t.Run("GET non-existing link", func(t *testing.T) {
		req, _ := http.NewRequest("GET", DefaultLocalhost + "/shorten/nonexisting", nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if !assert.NoError(t, err) {
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		if !assert.Equal(t, http.StatusBadRequest, "handler returned wrong status code") {
			return
		}
		//assert.Equal(t, "{\"This link is non-existing or expired\"}", w.Body.String())
	})
	// t.Run("POST empty body", func(t *testing.T) {
	// 	w := performRequest(http.HandlerFunc(postHandler), "POST", DefaultLocalhost + "/shorten/")

	// 	assert.Equal(t, http.StatusBadRequest, w.Code)
	// 	assert.Equal(t, "{\"Incorrect URL format\"}", w.Body.String())
	// })
}