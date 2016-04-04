package storage

import (
	"strings"
)

// Name takes a slash ('/') delimited package name and returns the name of the object to be stored in object storage
func Name(pkgName string) string {
	replaced := strings.Replace(pkgName, "/", "-", -1)
	return replaced + ".tar"
}
