package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/arschles/assert"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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

func isGoImportMeta(node *html.Node) bool {
	// if node.Type != html.ElementNode {
	// 	fmt.Printf("node is a %d, not an element\n", node.Type)
	// 	return false
	// }
	if node.DataAtom != atom.Meta {
		fmt.Printf("node is a %d atom, expected %d\n", node.DataAtom, atom.Meta)
		return false
	}
	return false
}

func TestGoGetExecute(t *testing.T) {
	const webHost = "goprox.io"
	const outwardPort = 80
	const gitScheme = "http"
	const gitHost = "git.goprox.io"
	type testCase struct {
		pkgName      string
		importPrefix string
		repoRoot     string
	}
	testCases := []testCase{
		testCase{
			pkgName:      "github.com/arschles/assert",
			importPrefix: "github.com/arschles/assert",
			repoRoot:     gitHost,
		},
	}

	hdl := goGet(webHost, outwardPort, gitScheme, gitHost)
	for i, tc := range testCases {
		r, err := http.NewRequest("GET", fmt.Sprintf("/%s?go-get=1", tc.pkgName), nil)
		if err != nil {
			t.Errorf("error creating request for case %d (%s)", i, err)
			continue
		}
		assert.NoErr(t, err)
		w := httptest.NewRecorder()
		hdl.ServeHTTP(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("expected status code was %d, got %d", http.StatusOK, w.Code)
			continue
		}
		htmlDoc, err := goquery.NewDocumentFromReader(w.Body)
		if err != nil {
			t.Errorf("error creating new HTML parser for case %d (%s)", i, err)
			continue
		}
		// var found *html.Node
		// var numFound = 0
		for _, node := range htmlDoc.Nodes {
			if isGoImportMeta(node) {
				// found = node
				// numFound++
			}
		}
		t.Skip("TODO: verify returned HTML")
		// if numFound != 1 {
		// 	t.Errorf("expected 1 go-import meta tag found for case %d, got %d", i, numFound)
		// 	continue
		// }
		// if found == nil {
		// 	t.Errorf("'go-import' meta tag not found for case %d", i)
		// 	continue
		// }
	}
}

func TestGoGetExecuteNonStandardPorts(t *testing.T) {
	t.Skip("TODO")
	t.SkipNow()
}

func TestStandardPort(t *testing.T) {
	assert.True(t, standardPort("http", 80), "80 was not reported a standard port for http")
	assert.True(t, standardPort("https", 443), "443 was not reported a standard port for https")
	assert.False(t, standardPort("http", 443), "443 was reported a standard port for http")
	assert.False(t, standardPort("https", 80), "80 was reported a standard port for https")
	assert.False(t, standardPort("http", 8080), "8080 was reported a standard port for http")
	assert.False(t, standardPort("https", 8081), "8081 was reported a standard port for https")
}

func TestGetImportPrefix(t *testing.T) {
	t.Skip("TODO")
}

func TestGetRepoRoot(t *testing.T) {
	t.Skip("TODO")
}

func TestGetPackage(t *testing.T) {
	type testCase struct {
		u        *url.URL
		expected string
	}
	testCases := []testCase{
		testCase{u: &url.URL{Path: "/a/b/c"}, expected: "a/b/c"},
	}
	for i, tc := range testCases {
		pkg := getPackage(tc.u)
		if tc.expected != pkg {
			t.Errorf("expected package %s for case %d, got %s", tc.expected, i, pkg)
			continue
		}
	}
}
