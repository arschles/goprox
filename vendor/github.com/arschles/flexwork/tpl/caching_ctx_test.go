package tpl

import (
	"testing"
)

func TestCachingCtxPrepare(t *testing.T) {
	errs := runCtxTest(NewCachingContext(getBaseDir(), funcs))
	if len(errs) > 0 {
		t.Fatalf("Errors: %s", errs)
	}
}
