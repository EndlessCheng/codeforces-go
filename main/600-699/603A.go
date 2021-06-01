package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// 纯 DP 推导见 https://www.luogu.com.cn/problem/solution/CF603A

// github.com/EndlessCheng/codeforces-go
func CF603A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s string
	Fscan(in, &n, &s)
	ans := strings.Count(s, "01") + strings.Count(s, "10") + 3
	if ans > n {
		ans = n
	}
	Fprint(out, ans)
}

//func main() { CF603A(os.Stdin, os.Stdout) }
