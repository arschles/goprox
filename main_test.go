package main

import (
	"testing"

	"github.com/arschles/assert"
	"github.com/arschles/flexwork/tpl"
)

func TestCreateTplCtx(t *testing.T) {
	tplDir := "rootfs/templates"
	ctx := createTplCtx(tplDir, funcs)
	tpl, err := ctx.Prepare(tpl.NewFiles("index.html"))
	assert.NoErr(t, err)
	assert.NotNil(t, tpl, "returned template")
}
