package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1515E(in io.Reader, out io.Writer) {
	var n int
	var mod, ans int64
	Fscan(in, &n, &mod)
	inv := make([]int64, n+1)
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (mod - mod/int64(i)) * inv[mod%int64(i)] % mod
	}

	dp := [2][][]int64{}
	for i := range dp {
		dp[i] = make([][]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, n+1)
		}
	}
	dp[1][1][1] = 1
	for i := 2; i <= n; i++ {
		cur, pre := dp[i&1], dp[i&1^1]
		for _, r := range cur {
			for j := range r {
				r[j] = 0
			}
		}
		cur[1][0] = pre[1][1]
		for j := 2; j <= i; j++ {
			for k := 1; k <= j; k++ {
				cur[j][0] += pre[j][k]
			}
			cur[j][0] %= mod
			cur[j][1] = pre[j-1][0] * int64(j) % mod
			for k := 2; k <= j; k++ {
				cur[j][k] = pre[j-1][k-1] * int64(j) * 2 % mod * inv[k] % mod
			}
		}
	}
	for _, r := range dp[n&1] {
		for _, v := range r[1:] {
			ans += v
		}
	}
	Fprint(out, ans%mod)
}

//func main() { CF1515E(os.Stdin, os.Stdout) }
