package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://oeis.org/A048004

// github.com/EndlessCheng/codeforces-go
func CF1027E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, lim int
	Fscan(in, &n, &lim)
	if lim == 1 {
		Fprint(out, 0)
		return
	}
	lim--
	k := min(n, lim)
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, k)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int64
	f = func(n, k int) int64 {
		if k < 0 || k > n {
			return 0
		}
		if k == 0 || k == n {
			return 1
		}
		dv := &dp[n][k]
		if *dv != -1 {
			return *dv
		}
		*dv = (f(n-1, k)*2 + f(n-1, k-1) - f(n-2, k-1)*2 + f(n-k-1, k-1) - f(n-k-2, k)) % mod
		return *dv
	}

	ans := int64(0)
	for mx := 1; mx <= k; mx++ {
		dp := make([]int64, n+1)
		dp[0] = 1
		for i := 1; i <= n; i++ {
			for j := i - 1; j >= 0 && (i-j)*mx <= lim; j-- {
				dp[i] += dp[j] // 这里可以用前缀和优化
			}
			dp[i] %= mod
		}
		ans += dp[n] * f(n-1, mx-1) % mod
	}
	Fprint(out, (ans*2%mod+mod)%mod)
}

//func main() { CF1027E(os.Stdin, os.Stdout) }
