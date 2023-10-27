package goutils

import (
	"github.com/virusdefender/goutils/assert"
	"testing"
)

func TestEscapeValidUtf8(t *testing.T) {
	origin := `\\xdb\xdb` + "\xbd\x20\xe2\x8c\x98你\u597d'\"\n\x00"
	result := EscapeToValidUtf8([]byte(origin))
	// fmt.Println(result)
	assert.Equal(t, result, `\\xdb\xdb\xbd ⌘你好'"`+"\n\\x00")
}
