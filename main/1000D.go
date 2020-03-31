package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1000D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 998244353
	const mx = 1000
	C := [mx + 1][mx + 1]int64{}
	for i := 0; i <= mx; i++ {
		C[i][0] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
		}
		C[i][i] = 1
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	dp := make([]int64, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if a[i] <= 0 {
			continue
		}
		for j := i + a[i] + 1; j <= n; j++ {
			dp[i] = (dp[i] + dp[j]*C[j-i-1][a[i]]) % mod
		}
	}
	ans := int64(0)
	for _, v := range dp[:n] {
		ans = (ans + v) % mod
	}
	Fprint(_w, ans)
}

//func main() { CF1000D(os.Stdin, os.Stdout) }
