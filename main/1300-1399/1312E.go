package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1312E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var f func(l, r int) int
	f = func(l, r int) int {
		if l == r {
			return a[l]
		}
		if dp[l][r] != 0 {
			return dp[l][r]
		}
		dp[l][r] = -1
		for m := l; m < r; m++ {
			if dpl := f(l, m); dpl > 0 && dpl == f(m+1, r) {
				dp[l][r] = dpl + 1
				break
			}
		}
		return dp[l][r]
	}

	dp2 := make([]int, n+1)
	for i := range dp2 {
		dp2[i] = 1e9
	}
	dp2[0] = 0
	for i := range dp2 {
		for j := i + 1; j <= n; j++ {
			if f(i, j-1) > 0 {
				if dp2[i]+1 < dp2[j] {
					dp2[j] = dp2[i] + 1
				}
			}
		}
	}
	Fprint(_w, dp2[n])
}

//func main() { CF1312E(os.Stdin, os.Stdout) }
