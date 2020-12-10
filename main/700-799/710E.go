package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF710E(in io.Reader, out io.Writer) {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	var n int
	var x, y int64
	Fscan(in, &n, &x, &y)
	dp := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		if i&1 > 0 {
			dp[i] = min(dp[i-1]+x, dp[i/2+1]+x+y)
		} else {
			dp[i] = min(dp[i-1]+x, dp[i/2]+y)
		}
	}
	Fprint(out, dp[n])
}

//func main() { CF710E(os.Stdin, os.Stdout) }
