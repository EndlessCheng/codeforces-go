package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF580D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, m, k, x, y int
	Fscan(in, &n, &m, &k)
	dp := make([][]int64, 1<<n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		dp[1<<i][i] = a[i]
	}
	ex := make([][]int64, n)
	for i := range ex {
		ex[i] = make([]int64, n)
	}
	for ; k > 0; k-- {
		Fscan(in, &x, &y)
		Fscan(in, &ex[x-1][y-1])
	}

	for s, ds := range dp {
		for S := uint(s); S > 0; S &= S - 1 {
			v := bits.TrailingZeros(S)
			for C := (1<<n - 1) &^ uint(s); C > 0; C &= C - 1 {
				w := bits.TrailingZeros(C)
				dp[s|1<<w][w] = max(dp[s|1<<w][w], ds[v]+ex[v][w]+a[w])
			}
		}
	}

	ans := int64(0)
	for s := 1<<m - 1; s < 1<<n; {
		for _, d := range dp[s] {
			ans = max(ans, d)
		}
		x := s & -s
		y := s + x
		s = s&^y/x>>1 | y
	}
	Fprint(out, ans)
}

//func main() { CF580D(os.Stdin, os.Stdout) }
