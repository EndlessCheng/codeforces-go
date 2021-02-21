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

// 替换常见变量名（数组、字符串等）
func renameInputArgs(funcDefineLine string) string {
	return strings.NewReplacer(
		"arr ", "a ",
		"nums ", "a ",
		"mat ", "a ",
		"matrix ", "a ",
		"grid ", "a ",
		"word ", "s ",
		"word1 ", "x ",
		"word2 ", "y ",
	).Replace(funcDefineLine)
}

func _parseReturnType(line string) string {
	i := strings.LastIndexByte(line, ')') + 2
	return line[i : len(line)-2]
}

func namedReturnFunc(name string) modifyLineFunc {
	return func(funcDefineLine string) string {
		returnType := _parseReturnType(funcDefineLine)
		if returnType == "" {
			return funcDefineLine
		} // 无返回值
		returnName := name
		if strings.HasPrefix(funcDefineLine, "func Constructor(") {
			returnName = "_"
		}
		i := strings.LastIndexByte(funcDefineLine, ')') + 2
		return funcDefineLine[:i] + "(" + returnName + " " + returnType + ") {"
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
			if _parseReturnType(lines[lo]) != "" {
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
