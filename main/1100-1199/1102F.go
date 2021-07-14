package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1102F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	minD := make([][]int, n)
	for i := range minD {
		minD[i] = make([]int, n)
		for j := range minD[i] {
			minD[i][j] = 1e9
		}
	}
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			for k := 0; k < i; k++ {
				minD[i][k] = min(minD[i][k], abs(a[i][j]-a[k][j]))
				minD[k][i] = minD[i][k]
			}
		}
	}
	if n == 1 {
		ans = 1e9
		for i, v := range a[0][:m-1] {
			ans = min(ans, abs(v-a[0][i+1]))
		}
		Fprint(out, ans)
		return
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for stI, stR := range a {
		for _, r := range dp {
			for j := range r {
				r[j] = 0
			}
		}
		for i, mi := range minD[stI] {
			if i != stI {
				dp[1<<i][i] = mi
			}
		}
		for s := 1; s < len(dp); s++ {
			for ss := uint(s); ss > 0; ss &= ss - 1 {
				i := bits.TrailingZeros(ss)
				for t, lb := 1<<n-1^s, 0; t > 0; t ^= lb {
					lb = t & -t
					if j := bits.TrailingZeros(uint(lb)); j != stI {
						dp[s|lb][j] = max(dp[s|lb][j], min(dp[s][i], minD[i][j]))
					}
				}
			}
		}
		for endI, dv := range dp[1<<n-1^(1<<stI)] {
			if endI != stI {
				for k, v := range a[endI][:m-1] {
					dv = min(dv, abs(v-stR[k+1]))
				}
				ans = max(ans, dv)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1102F(os.Stdin, os.Stdout) }
