package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1107E(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(in, &n, &s)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(l, r, ex int) int {
		if l > r {
			return 0
		}
		p := &dp[l][r][ex]
		if *p != -1 {
			return *p
		}
		res := f(l, r-1, 0) + a[ex+1]
		for k := l; k < r; k++ {
			if s[k] == s[r] {
				res = max(res, f(l, k, ex+1)+f(k+1, r-1, 0))
			}
		}
		*p = res
		return res
	}
	Fprint(out, f(0, n-1, 0))
}

//func main() { cf1107E(os.Stdin, os.Stdout) }
