package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF837D(in io.Reader, out io.Writer) {
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
	var n, k, ans int
	var v int64
	Fscan(in, &n, &k)
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, k*25+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for ; n > 0; n-- {
		Fscan(in, &v)
		c2 := bits.TrailingZeros64(uint64(v))
		c5 := 0
		for ; v%5 == 0; v /= 5 {
			c5++
		}
		for i := k; i > 0; i-- {
			for j := k * 25; j >= c5; j-- {
				if dp[i-1][j-c5] >= 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j-c5]+c2)
				}
			}
		}
	}
	for f, t := range dp[k] {
		ans = max(ans, min(f, t))
	}
	Fprint(out, ans)
}

//func main() { CF837D(os.Stdin, os.Stdout) }
