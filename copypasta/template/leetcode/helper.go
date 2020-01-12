package leetcode

import (
	"bytes"
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

func lowerFirstChar(codeS string) string {
	lower := func(c byte) byte {
		if 'A' <= c && c <= 'Z' {
			return c - 'A' + 'a'
		}
		return c
	}
	if !strings.HasPrefix(codeS, "func ") {
		// TODO
		return codeS
	}
	code := []byte(codeS)
	i := bytes.IndexByte(code, '(')
	code[i+1] = lower(code[i+1])
	for ; i < len(codeS); i++ {
		if code[i] == ',' {
			code[i+2] = lower(code[i+2])
		}
	}
	return string(code)
}

func namedReturn(code string, name string) string {
	lines := strings.Split(code, "\n")
	firstLine := lines[0]
	i := strings.Index(firstLine, ") ") + 2
	returnType := firstLine[i : len(firstLine)-2]
	lines[0] = firstLine[:i] + "(" + name + " " + returnType + ") {"
	return strings.Join(lines, "\n")
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
