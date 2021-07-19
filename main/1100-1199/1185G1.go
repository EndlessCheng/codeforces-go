package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1185G1(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var n, t, ans int
	Fscan(in, &n, &t)
	a := make([]int, n)
	tp := make([]int, n)
	for i := range a {
		Fscan(in, &a[i], &tp[i])
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range a {
		dp[1<<i][i] = 1
	}
	for s, dr := range dp {
		for ss := uint(s); ss > 0; ss &= ss - 1 {
			i := bits.TrailingZeros(ss)
			for t, lb := len(dp)-1^s, 0; t > 0; t ^= lb {
				lb = t & -t
				j := bits.TrailingZeros(uint(lb))
				if tp[j] != tp[i] {
					dp[s|lb][j] = (dp[s|lb][j] + dr[i]) % mod
				}
			}
		}
	}
	sum := make([]int, 1<<n)
	for i, v := range a {
		bit := 1 << i
		for mask := 0; mask < bit; mask++ {
			sum[bit|mask] = sum[mask] + v
		}
	}
	for i, s := range sum {
		if s == t {
			for ss := uint(i); ss > 0; ss &= ss - 1 {
				ans = (ans + dp[i][bits.TrailingZeros(ss)]) % mod
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1185G1(os.Stdin, os.Stdout) }
