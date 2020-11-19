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
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int32, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	dp := make([][][2]int32, n+1)
	for i := range dp {
		dp[i] = make([][2]int32, m+1)
		for j := range dp[i] {
			dp[i][j] = [2]int32{-1e9, -1e9}
		}
	}
	dp[1][0][0] = 0
	dp[1][1][1] = 0
	for i := 2; i <= n; i++ {
		dp[i][0][0] = dp[i-1][0][0]
		for j := 1; j <= m; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1])
			dp[i][j][1] = max(dp[i-1][j-1][0], dp[i-1][j-1][1]+a[i])
		}
	}
	ans := max(dp[n][m][0], dp[n][m][1])
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = [2]int32{-1e9, -1e9}
		}
	}
	dp[1][0][0] = 0
	dp[1][1][1] = a[1]
	for i := 2; i <= n; i++ {
		dp[i][0][0] = dp[i-1][0][0]
		for j := 1; j <= m; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1])
			dp[i][j][1] = max(dp[i-1][j-1][0], dp[i-1][j-1][1]+a[i])
		}
	}
	ans = max(ans, dp[n][m][1])
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
