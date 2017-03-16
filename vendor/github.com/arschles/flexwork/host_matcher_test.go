package flexwork

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arschles/assert"
)

func TestMatchHost(t *testing.T) {
	hdl := MatchHost(map[string]http.Handler{
		"somehost.com": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	})
	srv := httptest.NewServer(hdl)
	defer srv.Close()

	r1, err := http.NewRequest("GET", srv.URL, nil)
	assert.NoErr(t, err)
	r1.Host = "someotherhost.com"
	res1, err := http.DefaultClient.Do(r1)
	assert.NoErr(t, err)
	assert.Equal(t, res1.StatusCode, http.StatusNotFound, "response code")

	r2, err := http.NewRequest("GET", srv.URL, nil)
	assert.NoErr(t, err)
	r2.Host = "somehost.com"
	res2, err := http.DefaultClient.Do(r2)
	assert.NoErr(t, err)
	assert.Equal(t, res2.StatusCode, http.StatusOK, "response code")
}
