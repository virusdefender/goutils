package goutils

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// 将非 utf8 字符和部分控制字符转换为 `\x` 的模式，函数返回的结果和原始字符串是非等价的
// 一般用于控制台打印等场景，避免出现乱码
func EscapeToValidUtf8(s []byte) string {
	ret := make([]rune, 0, len(s))
	start := 0
	for {
		r, size := utf8.DecodeRune(s[start:])
		if r == utf8.RuneError {
			// 说明是空的
			if size == 0 {
				break
			} else {
				// 不是 rune
				ret = append(ret, []rune(fmt.Sprintf("\\x%02x", s[start]))...)
			}
		} else {
			// 不是换行之类的控制字符
			if unicode.IsControl(r) && !unicode.IsSpace(r) {
				ret = append(ret, []rune(fmt.Sprintf("\\x%02x", r))...)
			} else {
				// 正常字符
				ret = append(ret, r)
			}
		}
		start += size
	}
	return string(ret)
}
