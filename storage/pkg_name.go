package storage

import (
	"strings"
)

const (
	dotTar = ".tar"
)

// Name takes a slash ('/') delimited package name and returns the name of the object to be stored in object storage
func Name(pkgName string) string {
	replaced := strings.Replace(pkgName, "/", "-", -1)
	return replaced + dotTar
}

// ReverseName takes a hyphen ('-') deliminted package name with a '.tar' on the end and returns the standard slash ('/') deliminted package name without the '.tar' on the end
func ReverseName(storageName string) string {
	lastIdx := strings.LastIndex(storageName, dotTar)
	if lastIdx == -1 {
		return ""
	}
	noTar := storageName[:lastIdx]
	return strings.Replace(noTar, "-", "/", -1)
}
