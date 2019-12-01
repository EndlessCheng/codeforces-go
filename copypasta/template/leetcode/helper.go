package leetcode

import (
	"strings"
	"unicode"
)

func parseFuncName(code string) (string, bool) {
	code = strings.TrimSpace(code)
	if !strings.HasPrefix(code, "func ") {
		return "", false
	}
	i := strings.IndexByte(code, '(')
	return code[5:i], true
}

func isASCII(s string) bool {
	// this is faster than `for _, c := range s`
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
