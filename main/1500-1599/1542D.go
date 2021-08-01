package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1542D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n int
	var op string
	Fscan(in, &n)
	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		if Fscan(in, &op); op == "+" {
			Fscan(in, &a[i])
		}
	}

	ans := int64(0)
	for p := 1; p <= n; p++ {
		v := a[p]
		if v == 0 {
			continue
		}
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, n+2)
		}
		dp[0][0] = 1
		for i := 1; i <= n; i++ {
			w := a[i]
			for j := 0; j <= n; j++ {
				if w == 0 {
					if j > 0 || i <= p {
						dp[i][max(j-1, 0)] = (dp[i][max(j-1, 0)] + dp[i-1][j]) % mod
					}
				} else if w < v || w == v && i < p {
					dp[i][j+1] = (dp[i][j+1] + dp[i-1][j]) % mod
				} else {
					dp[i][j] = (dp[i][j] + dp[i-1][j]) % mod
				}
				if i != p {
					dp[i][j] = (dp[i][j] + dp[i-1][j]) % mod
				}
			}
		}
		s := int64(0)
		for _, fv := range dp[n] {
			s += int64(fv)
		}
		ans = (ans + s%mod*v) % mod
	}
	Fprint(out, ans)
}

//func main() { CF1542D(os.Stdin, os.Stdout) }
