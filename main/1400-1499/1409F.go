package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1409F(in io.Reader, out io.Writer) {
	var n, K int
	var s, t string
	Fscan(in, &n, &K, &s, &t)
	if t[0] == t[1] {
		c := strings.Count(s, t[:1]) + K
		if c > n {
			c = n
		}
		Fprint(out, c*(c-1)/2)
		return
	}
	I := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, K+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(p, left, c0 int) int
	f = func(p, left, c0 int) (res int) {
		if p == n {
			return
		}
		dv := &dp[p][left][c0]
		if *dv >= 0 {
			return *dv
		}
		defer func() { *dv = res }()
		res = I(s[p] == t[1])*c0 + f(p+1, left, c0+I(s[p] == t[0])) // 不修改
		if left > 0 {
			res = max(res, f(p+1, left-1, c0+1))  // 修改成 t[0]
			res = max(res, c0+f(p+1, left-1, c0)) // 修改成 t[1]
		}
		return
	}
	Fprint(out, f(0, K, 0))
}

//func main() { CF1409F(os.Stdin, os.Stdout) }
