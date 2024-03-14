package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF906C(in io.Reader, out io.Writer) {
	var n, m, v, w int
	Fscan(in, &n, &m)
	if m == n*(n-1)/2 {
		Fprint(out, 0)
		return
	}

	g := make([]int, n)
	for i := range g {
		g[i] |= 1 << i // 自己认识自己
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] |= 1 << w
		g[w] |= 1 << v
	}

	dp := make([]int, 1<<n)
	use := make([]int, 1<<n)
	from := make([]int, 1<<n)
	for i := range dp {
		dp[i] = 1e9
	}
	for i, mask := range g {
		dp[mask] = 1
		use[mask] = i
	}
	for s, dv := range dp {
		if dv == 1e9 { // 巨大优化
			continue
		}
		for t := uint(s); t > 0; t &= t - 1 {
			i := bits.TrailingZeros(t)
			if next := s | g[i]; dv+1 < dp[next] {
				dp[next] = dv + 1
				use[next] = i
				from[next] = s
			}
		}
	}

	Fprintln(out, dp[1<<n-1])
	for i := 1<<n - 1; i > 0; i = from[i] {
		Fprint(out, use[i]+1, " ") // 任意顺序均可
	}
}

//func main() { CF906C(os.Stdin, os.Stdout) }
