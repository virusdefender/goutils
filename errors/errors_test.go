package errors

import (
	"errors"
	"github.com/virusdefender/goutils/assert"
	"testing"
)

func TestErrors(t *testing.T) {
	err := errors.New("error")
	wrapped := wrap(err, "call api1")
	assert.ErrorStringEqual(t, wrapped, "call api1: error")
	assert.ErrorStringEqual(t, errors.Unwrap(wrapped), "error")
}
