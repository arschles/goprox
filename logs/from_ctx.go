package logs

import (
	"context"
	"io/ioutil"
	"log"
	"os"
)

const (
	ctxDebugVarName = "debug"
)

// FromContext returns a new logger from the given context
func FromContext(ctx context.Context) *log.Logger {
	if ctx.Value(ctxDebugVarName) == true {
		return log.New(os.Stdout, "[Debug] ", log.Lshortfile)
	}
	return log.New(ioutil.Discard, "", 0)
}

// DebugContext returns a child context from ctx that indicates that debug
// logs should be turned on
func DebugContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxDebugVarName, true)
}
