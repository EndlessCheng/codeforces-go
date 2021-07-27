package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1215E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var n, v int
	cnt := make([]int, 20)
	inv := [20][20]int64{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		v--
		for i, c := range cnt {
			inv[v][i] += int64(c)
		}
		cnt[v]++
	}

	const m = 1 << 20
	dp := make([]int64, m)
	for i := range dp {
		dp[i] = 1e18
	}
	dp[0] = 0
	for s, dv := range dp {
		for t, lb := m-1^s, 0; t > 0; t ^= lb {
			lb = t & -t
			p := bits.TrailingZeros(uint(lb))
			sum := dv
			for ss := uint(s); ss > 0; ss &= ss - 1 {
				q := bits.TrailingZeros(ss)
				sum += inv[p][q]
			}
			dp[s|lb] = min(dp[s|lb], sum)
		}
	}
	Fprint(out, dp[m-1])
}

//func main() { CF1215E(os.Stdin, os.Stdout) }
