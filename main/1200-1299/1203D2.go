package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1203D2(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)
	suf := make([]int, m+1)
	suf[m] = n
	for i, j := n-1, m-1; j >= 0; i-- {
		if s[i] == t[j] {
			suf[j] = i
			j--
		}
	}
	ans := suf[0]
	for i, j := 0, 0; j < m; i++ {
		if s[i] == t[j] {
			j++
			ans = max(ans, suf[j]-i-1)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1203D2(os.Stdin, os.Stdout) }
