package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1467D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7

	var n, k, q, p int
	Fscan(in, &n, &k, &q)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, k+1)
		dp[i][0] = 1
	}
	for j := 1; j <= k; j++ {
		dp[0][j] = dp[1][j-1]
		for i := 1; i < n-1; i++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i+1][j-1]) % mod
		}
		dp[n-1][j] = dp[n-2][j-1]
	}
	w := make([]int64, n)
	for i, d := range dp {
		for j := 0; j <= k; j++ {
			w[i] += d[j] * d[k-j] % mod
		}
		w[i] %= mod
	}

	var s, v int64
	a := make([]int64, n)
	for i, w := range w {
		Fscan(in, &a[i])
		s += w * a[i] % mod
	}
	for ; q > 0; q-- {
		Fscan(in, &p, &v)
		p--
		s = ((s+(v-a[p])*w[p])%mod + mod) % mod
		a[p] = v
		Fprintln(out, s)
	}
}

//func main() { CF1467D(os.Stdin, os.Stdout) }
