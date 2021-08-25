package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1012C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
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
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	m := (n + 1) / 2

	// dp[i][j][0/1] 表示前 i 个元素组成 j 个山峰的答案，0 表示不修改 a[i]，1 表示修改 a[i] 使其能够让 a[i+1] 成为山峰
	// 转移的时候只需考虑是否修改当前元素，不考虑是否让当前元素成为山峰
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, m+1)
		for j := 1; j <= m; j++ {
			dp[i][j] = [2]int{1e9, 1e9}
		}
	}
	for i := 1; i <= n; i++ {
		cost := max(a[i]-a[i+1]+1, 0)
		dp[i][0][1] = cost
		for j := 1; j <= m; j++ {
			dp[i][j][0] = dp[i-1][j][0]
			dp[i][j][1] = dp[i-1][j][0] + cost
			if i > 1 {
				dp[i][j][0] = min(dp[i][j][0], dp[i-2][j-1][1]+max(a[i]-a[i-1]+1, 0))
				dp[i][j][1] = min(dp[i][j][1], dp[i-2][j-1][1]+max(a[i]-min(a[i-1], a[i+1])+1, 0))
			}
		}
	}
	for k := 1; k <= m; k++ {
		Fprint(out, min(dp[n][k][0], dp[n-1][k-1][1]), " ")
	}
}

//func main() { CF1012C(os.Stdin, os.Stdout) }
