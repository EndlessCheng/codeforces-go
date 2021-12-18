package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	t := []byte(s)
	x := 0
	for i, c := range s {
		if c == '(' {
			x++
		} else if x > 0 {
			x--
		} else {
			t[i] = ' ' // 从左往右扫描，将无法与左括号匹配的右括号标记为空格
		}
	}

	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			x++
		} else if x > 0 {
			x--
		} else {
			t[i] = ' ' // 从右往左扫描，将无法与右括号匹配的左括号标记为空格
		}
	}

	// 按空格划分字符串，统计最长的子串及该长度子串出现次数
	maxL, cnt := 0, 1
	for _, s := range bytes.Fields(t) {
		if len(s) > maxL {
			maxL, cnt = len(s), 1
		} else if len(s) == maxL {
			cnt++
		}
	}
	Fprint(out, maxL, cnt)
}

func main() { run(os.Stdin, os.Stdout) }
