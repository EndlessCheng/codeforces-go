package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1580B(in io.Reader, out io.Writer) {
	var n, tarDep, k, mod int
	Fscan(in, &n, &tarDep, &k, &mod)
	const mx = 100
	F := [mx]int{1}
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	C := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	dp := make([][][]int, tarDep)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for p := range dp[i][j] {
				dp[i][j][p] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(dep, size, need int) (res int) {
		if dep < 0 {
			if need > 0 {
				return
			}
			return F[size]
		}
		if size == 0 {
			return 1
		}
		// 无明显优化
		// if (need-1)>>dep > 0 { return }
		p := &dp[dep][size][need]
		if *p >= 0 {
			return *p
		}
		if dep == 0 {
			need--
		}
		for leftSz := range size {
			for leftNeed := max(need-(size-1-leftSz), 0); leftNeed <= min(leftSz, need); leftNeed++ {
				leftRes := f(dep-1, leftSz, leftNeed)
				if leftRes == 0 { // 可以快一倍
					continue
				}
				res = (res + C[size-1][leftSz]*leftRes%mod*f(dep-1, size-1-leftSz, need-leftNeed)) % mod
			}
		}
		*p = res
		return
	}
	Fprint(out, f(tarDep-1, n, k))
}

//func main() { cf1580B(os.Stdin, os.Stdout) }
