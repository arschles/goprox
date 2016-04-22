package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arschles/assert"
)

func TestMatchHost(t *testing.T) {
	hdl := MatchHost(map[string]http.Handler{
		"somehost.com": http.RedirectHandler("abc.com", http.StatusFound),
	})

	r1, err := http.NewRequest("GET", "someotherhost.com", nil)
	assert.NoErr(t, err)
	r1.Header.Set("Host", "someotherhost.com")
	w1 := httptest.NewRecorder()
	hdl.ServeHTTP(w1, r1)
	assert.Equal(t, w1.Code, http.StatusNotFound, "response code")

	r2, err := http.NewRequest("GET", "somehost.com", nil)
	assert.NoErr(t, err)
	r2.Header.Set("Host", "somehost.com")
	w2 := httptest.NewRecorder()
	hdl.ServeHTTP(w2, r2)
	assert.Equal(t, w2.Code, http.StatusFound, "response code")
}
