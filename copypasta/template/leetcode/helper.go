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
			if strings.HasPrefix(line, "func Constructor(") { // 构造器定义
				funcLos = append(funcLos, lo)
			} else if strings.HasPrefix(line, "func (") { // 方法定义
				funcLos = append(funcLos, lo)
			}
		}
		funcName = "Constructor"
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

func _toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func toGolangReceiverName(funcDefineLine string) string {
	if !strings.HasPrefix(funcDefineLine, "func (this *") {
		return funcDefineLine
	}
	// 由于采用的是全局变量的写法，receiver 可以去掉
	return strings.Replace(funcDefineLine, "this *", "", 1)
}

func lowerArgsFirstChar(funcDefineLine string) string {
	code := []byte(funcDefineLine)
	i := bytes.LastIndexByte(code, '(')
	code[i+1] = _toLower(code[i+1])
	for ; i < len(code); i++ {
		if code[i] == ',' {
			code[i+2] = _toLower(code[i+2])
		}
	}
	return string(code)
}

func _parseReturnType(line string) string {
	i := strings.LastIndexByte(line, ')')
	return strings.TrimSpace(line[i+1 : len(line)-2])
}

func namedReturnFunc(name string) modifyLineFunc {
	return func(funcDefineLine string) string {
		returnType := _parseReturnType(funcDefineLine)
		if returnType == "" {
			return funcDefineLine
		} // 无返回值
		if returnType == "int64" {
			return funcDefineLine
		}
		returnName := name
		if strings.HasPrefix(funcDefineLine, "func Constructor(") {
			returnName = "_"
		}
		i := strings.LastIndexByte(funcDefineLine, ')') + 2
		return funcDefineLine[:i] + "(" + returnName + " " + returnType + ") {"
	}
}

func modifyDefaultCode(code string, funcLos []int, funcList []modifyLineFunc, customFuncContent string) string {
	sep := "\n"
	if strings.ContainsRune(code, '\r') {
		sep = "\r\n"
	}
	lines := strings.Split(code, sep)
	for _, lo := range funcLos {
		if tp := _parseReturnType(lines[lo]); tp != "" {
			if tp == "int64" {
				customFuncContent = "\tans64 := func() (ans int) {\n\t\t\n\t\treturn\n\t}()" + customFuncContent /* return */ + " int64(ans64)"
			}
			lines[lo+1] = customFuncContent
		}
		for _, f := range funcList {
			lines[lo] = f(lines[lo])
		}
	}
	return strings.Join(lines, sep)
}

//

func findNonASCII(s string) int {
	for i, c := range s {
		if c > unicode.MaxASCII {
			return i // [0,i-1] 中的都是 ASCII
		}
	}
	return -1
}

func absPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
