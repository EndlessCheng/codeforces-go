package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// 注：O(n) 解法 https://www.luogu.com.cn/problem/solution/CF1109B

// github.com/EndlessCheng/codeforces-go
func CF1109B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s string
	Fscan(in, &s)
	n := len(s)
	if strings.Count(s[:n/2], string(s[0])) == n/2 {
		Fprint(out, "Impossible")
		return
	}
o:
	for i := 1; i < n; i++ {
		t := s[i:] + s[:i]
		if t == s {
			continue
		}
		for j := 0; j < n/2; j++ {
			if t[j] != t[n-1-j] {
				continue o
			}
		}
		Fprint(out, 1)
		return
	}
	Fprint(out, 2)
}

//func main() { CF1109B(os.Stdin, os.Stdout) }
