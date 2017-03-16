package examples

import (
	"errors"
	"github.com/arschles/assert"
	"testing"
)

func TestNil(t *testing.T) {
	assert.Nil(t, nil, "nil")
	assert.NotNil(t, "abc", "string")
}

func TestBooleans(t *testing.T) {
	assert.True(t, true, "boolean true")
	assert.False(t, false, "boolean false")
}

func TestEqual(t *testing.T) {
	s1 := struct {
		a string
		b int
	}{"testString", 1}
	s2 := struct {
		a string
		b int
	}{"testString", 1}
	assert.Equal(t, s1, s2, "anonymous struct")
}

func TestErrors(t *testing.T) {
	err1 := errors.New("this is an error")
	var err2 error
	assert.Err(t, err1, errors.New("this is an error"))
	assert.NoErr(t, err2)
	assert.ExistsErr(t, err1, "valid error")
}

func TestWithCustomAssertion(t *testing.T) {
	assertCustom(t, "foobar", "foobar")
}

func assertCustom(t *testing.T, s1, s2 string) {
	// When creating custom assertions, use assert.WithFrameWrapper to wrap t. Pass the wrapped t
	// to other assertions. This helps the assert library rewind the callstack to the appropriate
	// point when displaying the source of a failed assertion.
	wt := assert.WithFrameWrapper(t)
	assert.Equal(wt, s1, s2, "sample string")
}
