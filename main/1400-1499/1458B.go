package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1458B(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, cap, cur, totWater int
	Fscan(in, &n)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n*100+1)
		for j := range dp[i] {
			dp[i][j] = -1e9
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		Fscan(in, &cap, &cur)
		totWater += cur
		for j := i; j > 0; j-- {
			for k := j * 100; k >= cap; k-- {
				dp[j][k] = max(dp[j][k], dp[j-1][k-cap]+cur)
			}
		}
	}
	for _, d := range dp[1:] {
		mx := 0
		for cap, maxWater := range d {
			mx = max(mx, min(cap*2, maxWater+totWater))
		}
		Fprintf(out, "%.1f ", float64(mx)/2)
	}
}

//func main() { CF1458B(os.Stdin, os.Stdout) }
