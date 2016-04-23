package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arschles/assert"
)

func TestLogAndErr(t *testing.T) {
	w := httptest.NewRecorder()
	logAndErr(w, http.StatusNotFound, "abc%s", "d")
	assert.Equal(t, w.Code, http.StatusNotFound, "response code")
	assert.Equal(t, strings.TrimSpace(string(w.Body.Bytes())), "abcd", "response body")
}
