package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1203D2(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	for i, j := n-1, m-1; i > 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = m - 1 - j
	}
	ans := 0
	for l, r, j := 0, 1, 0; r <= n; r++ {
		for j+suf[r] < m {
			if s[l] == t[j] {
				j++
			}
			l++
		}
		if r-l > ans {
			ans = r - l
		}
	}
	Fprint(out, ans)
}

//func main() { CF1203D2(os.Stdin, os.Stdout) }
