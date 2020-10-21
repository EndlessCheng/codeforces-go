package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF682D(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var n, m, K int
	var s, t []byte
	Fscan(bufio.NewReader(in), &n, &m, &K, &s, &t)
	dp := make([][][][2]int, n)
	for i := range dp {
		dp[i] = make([][][2]int, m)
		for j := range dp[i] {
			dp[i][j] = make([][2]int, K+1)
			for k := range dp[i][j] {
				dp[i][j][k] = [2]int{-1, -1}
			}
		}
	}
	var f func(i, j, k, prevSame int) int
	f = func(i, j, k, prevSame int) (res int) {
		if k > K {
			return -1e9
		}
		if i == n || j == m {
			return
		}
		dv := &dp[i][j][k][prevSame]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = max(f(i+1, j, k, 0), f(i, j+1, k, 0))
		if s[i] == t[j] {
			res = max(res, 1+f(i+1, j+1, k+1-prevSame, 1))
		}
		return
	}
	Fprint(out, f(0, 0, 0, 0))
}

//func main() { CF682D(os.Stdin, os.Stdout) }
