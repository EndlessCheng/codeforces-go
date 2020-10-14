package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1132F(_r io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(bufio.NewReader(_r), &n, &s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var f func(int, int) int
	f = func(l, r int) (res int) {
		if l > r {
			return 0
		}
		if l == r {
			return 1
		}
		dv := &dp[l][r]
		if *dv > 0 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1 + f(l+1, r)
		for i := l + 1; i <= r; i++ {
			if s[i] == s[l] {
				if r := f(l+1, i-1) + f(i, r); r < res {
					res = r
				}
			}
		}
		return
	}
	Fprint(out, f(0, n-1))
}

//func main() { CF1132F(os.Stdin, os.Stdout) }
