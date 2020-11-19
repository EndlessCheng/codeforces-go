package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e9

	var n, m int
	Fscan(in, &m, &n)
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
			for k := range dp[i][j] {
				dp[i][j][k] = inf
			}
		}
	}
	dp[0][1][2] = 0
	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &pos[i])
		pos[i]--
		p0, p := pos[i-1], pos[i]
		for j, dj := range dp[i-1] {
			for k, dv := range dj {
				if dv == inf {
					continue
				}
				if j != p && k != p {
					dp[i][j][k] = min(dp[i][j][k], dv+a[p0][p])
				}
				if p0 != p && k != p {
					dp[i][p0][k] = min(dp[i][p0][k], dv+a[j][p])
				}
				if p0 != p && j != p {
					dp[i][j][p0] = min(dp[i][j][p0], dv+a[k][p])
				}
			}
		}
	}
	ans := inf
	for _, row := range dp[n] {
		for _, v := range row {
			ans = min(ans, v)
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
