package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1198D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := make([][50][50][50]int, n)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}
	var f func(x1, y1, x2, y2 int) int
	f = func(x1, y1, x2, y2 int) (res int) {
		if x1 == x2 && y1 == y2 {
			if a[x1][y1] == '#' {
				return 1
			}
			return
		}
		dv := &dp[x1][y1][x2][y2]
		if *dv != -1 {
			return *dv
		}
		res = max(x2-x1+1, y2-y1+1)
		for i := x1; i < x2; i++ {
			res = min(res, f(x1, y1, i, y2)+f(i+1, y1, x2, y2))
		}
		for i := y1; i < y2; i++ {
			res = min(res, f(x1, y1, x2, i)+f(x1, i+1, x2, y2))
		}
		*dv = res
		return
	}
	Fprint(out, f(0, 0, n-1, n-1))
}

//func main() { CF1198D(os.Stdin, os.Stdout) }
