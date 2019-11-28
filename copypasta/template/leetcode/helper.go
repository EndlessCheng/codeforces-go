package leetcode

import "strings"

func parseFuncName(code string) string {
	code = strings.TrimSpace(code)
	i := strings.IndexByte(code, '(')
	return code[5:i]
}
