package leetcode

import "strings"

func parseFuncName(code string) (string, bool) {
	code = strings.TrimSpace(code)
	if !strings.HasPrefix(code, "func ") {
		return "", false
	}
	i := strings.IndexByte(code, '(')
	return code[5:i], true
}
