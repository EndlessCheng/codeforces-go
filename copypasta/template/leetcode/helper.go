package leetcode

import (
	"strings"
	"unicode"
)

func parseFuncName(code string) (funcName string, isFuncProblem bool) {
	code = strings.TrimSpace(code)
	if !strings.HasPrefix(code, "func ") {
		return
	}
	i := strings.IndexByte(code, '(')
	return code[5:i], true
}

func findASCII(s string) int {
	// this is faster than `for _, c := range s`, because there is no rune conversion
	for i := range s {
		if s[i] > unicode.MaxASCII {
			return i
		}
	}
	return -1
}
