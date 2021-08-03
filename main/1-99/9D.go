package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF9D(in io.Reader, out io.Writer) {
	var N, H int
	Fscan(in, &N, &H)

	// dp[n][h] 表示表示由 n 个点组成的高度不超过 h 的二叉树的个数
	dp := make([][]int64, N+1)
	for i := range dp {
		dp[i] = make([]int64, N+1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for h := 1; h <= N; h++ {
		for n := 1; n <= N; n++ {
			for left := 0; left < n; left++ {
				dp[n][h] += dp[left][h-1] * dp[n-1-left][h-1]
			}
		}
	}
	Fprint(out, dp[N][N]-dp[N][H-1])
}

//func main() { CF9D(os.Stdin, os.Stdout) }
