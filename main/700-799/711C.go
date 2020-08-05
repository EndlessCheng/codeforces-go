package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF711C(_r io.Reader, out io.Writer) {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n, m, tar int
	Fscan(in, &n, &m, &tar)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	cost := make([][]int64, n)
	for i := range cost {
		cost[i] = make([]int64, m)
		for j := range cost[i] {
			Fscan(in, &cost[i][j])
		}
	}

	const inf int64 = 1e18
	ans := inf
	dp := make([][][]int64, n)
	for i := range dp {
		dp[i] = make([][]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, m)
			for k := range dp[i][j] {
				dp[i][j][k] = inf
			}
		}
	}
	if c := a[0]; c > 0 {
		dp[0][1][c-1] = 0
	} else {
		dp[0][1] = cost[0]
	}
	for i := 1; i < n; i++ {
		for b := 1; b <= n; b++ {
			for c, cst := range cost[i] {
				if a[i] > 0 {
					if c != a[i]-1 {
						continue
					}
					cst = 0
				}
				for pc := 0; pc < m; pc++ {
					bb := b
					if c != pc {
						bb--
					}
					dp[i][b][c] = min(dp[i][b][c], dp[i-1][bb][pc]+cst)
				}
			}
		}
	}
	for _, v := range dp[n-1][tar] {
		ans = min(ans, v)
	}
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF711C(os.Stdin, os.Stdout) }
