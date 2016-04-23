package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arschles/assert"
)

func TestGoGetNoQueryString(t *testing.T) {
	hdl := goGet("a", 1, "b", "c")
	r, err := http.NewRequest("GET", "/a/b/c", nil)
	assert.NoErr(t, err)
	w := httptest.NewRecorder()
	hdl.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest, "response code")
	assert.Equal(t, strings.TrimSpace(string(w.Body.Bytes())), `?go-get="1" expected`, "response body")
}
