package leetcode

import (
	"bytes"
	"path/filepath"
	"strings"
	"unicode"
)

// 细分的话有四种题目类型：
// 函数-无预定义类型，绝大多数题目都是这个类型
// 函数-有预定义类型，如 LC174C
// 方法-无预定义类型，如 LC175C
// 方法-有预定义类型，如 LC163B
func parseCode(code string) (funcName string, isFuncProblem bool, funcLos []int) {
	lines := strings.Split(code, "\n")
	if strings.Contains(code, "func Constructor(") {
		// 编写方法
		for lo, line := range lines {
			if strings.HasPrefix(line, "func (") { // 方法定义
				funcLos = append(funcLos, lo)
			}
		}
	} else {
		// 编写函数
		for lo, line := range lines {
			if strings.HasPrefix(line, "func ") { // 函数定义
				i := strings.IndexByte(line, '(')
				return strings.TrimSpace(line[5:i]), true, []int{lo}
			}
		}
	}
	return
}

//

type modifyLineFunc func(funcDefineLine string) string

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func toGolangReceiverName(funcDefineLine string) string {
	if !strings.HasPrefix(funcDefineLine, "func (this *") {
		return funcDefineLine
	}
	receiverName := ""
	for _, r := range funcDefineLine {
		if r == ')' {
			break
		}
		if unicode.IsUpper(r) {
			receiverName += string(toLower(byte(r)))
		}
	}
	return "func (" + receiverName + funcDefineLine[10:]
}

func lowerArgsFirstChar(funcDefineLine string) string {
	code := []byte(funcDefineLine)
	i := bytes.LastIndexByte(code, '(')
	code[i+1] = toLower(code[i+1])
	for ; i < len(code); i++ {
		if code[i] == ',' {
			code[i+2] = toLower(code[i+2])
		}
	}
	return string(code)
}

func parseReturnType(line string) string {
	i := strings.LastIndexByte(line, ')') + 2
	return line[i : len(line)-2]
}

func namedReturnFunc(name string) modifyLineFunc {
	return func(funcDefineLine string) string {
		returnType := parseReturnType(funcDefineLine)
		if returnType == "" {
			return funcDefineLine
		}
		i := strings.LastIndexByte(funcDefineLine, ')') + 2
		return funcDefineLine[:i] + "(" + name + " " + returnType + ") {"
	}
}

func modifyDefaultCode(code string, funcLos []int, fs []modifyLineFunc, customFuncContent string) string {
	sep := "\n"
	if strings.ContainsRune(code, '\r') {
		sep = "\r\n"
	}
	lines := strings.Split(code, sep)
	for _, lo := range funcLos {
		for _, f := range fs {
			if parseReturnType(lines[lo]) != "" {
				lines[lo+1] = customFuncContent
			}
			lines[lo] = f(lines[lo])
		}
	}
	return strings.Join(lines, sep)
}

//

func findNonASCII(s string) int {
	// this is faster than `for _, c := range s`, because there is no rune conversion
	for i := range s {
		if s[i] > unicode.MaxASCII {
			return i
		}
	}
	return -1
}

func absPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
