package tpl

import (
	"html/template"
)

// Context is a group of template files that all live under a directory
type Context interface {
	// Prepare renders all of the given Files for execution. Some
	// implementations may return a cached executor for the given Files
	Prepare(Files) (*template.Template, error)
}
