package goutils

import (
	"github.com/virusdefender/goutils/assert"
	"os"
	"testing"
)

func TestEnvOrDefault(t *testing.T) {
	err := os.Setenv("TEST_STRING_ENV1", "test")
	assert.Nil(t, err)
	err = os.Setenv("TEST_INT_ENV1", "1")

	assert.Equal(t, EnvOrDefaultString("TEST_STRING_ENV", "default"), "default")
	assert.Equal(t, EnvOrDefaultString("TEST_STRING_ENV1", "default"), "test")

	assert.Equal(t, EnvOrDefaultNumber("TEST_INT_ENV", 2), 2)
	assert.Equal(t, EnvOrDefaultNumber("TEST_INT_ENV1", 1), 1)
	assert.Equal(t, EnvOrDefaultNumber("TEST_STRING_ENV", 3), 3)
}
