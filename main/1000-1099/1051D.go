package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1051D(_r io.Reader, _w io.Writer) {
	const m = 998244353
	var n, k int
	Fscan(_r, &n, &k)
	if k == 1 {
		Fprint(_w, 2)
		return
	}
	dp := make([][][4]uint, n)
	for i := range dp {
		dp[i] = make([][4]uint, k+1)
	}
	dp[0][1] = [4]uint{1, 0, 0, 1}
	dp[0][2] = [4]uint{0, 1, 1, 0}
	for i := 1; i < n; i++ {
		dp[i][1][0] = dp[i-1][1][0]
		dp[i][1][3] = dp[i-1][1][3]
		for j := 2; j <= k; j++ {
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1] + dp[i-1][j][2] + dp[i-1][j-1][3]) % m
			dp[i][j][1] = (dp[i-1][j-1][0] + dp[i-1][j][1] + dp[i-1][j-2][2] + dp[i-1][j-1][3]) % m
			dp[i][j][2] = (dp[i-1][j-1][0] + dp[i-1][j-2][1] + dp[i-1][j][2] + dp[i-1][j-1][3]) % m
			dp[i][j][3] = (dp[i-1][j-1][0] + dp[i-1][j][1] + dp[i-1][j][2] + dp[i-1][j][3]) % m
		}
	}
	Fprint(_w, (dp[n-1][k][0]+dp[n-1][k][1]+dp[n-1][k][2]+dp[n-1][k][3])%m)
}

//func main() { CF1051D(os.Stdin, os.Stdout) }
