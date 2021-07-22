package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1120C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y int
	var s []byte
	Fscan(in, &n, &x, &y, &s)
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

	lcs := make([][]int, n+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}
	for i, v := range s {
		for j, w := range s {
			if v == w {
				lcs[i+1][j+1] = lcs[i][j] + 1
			}
		}
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + x
		for j := 1; j < i; j++ {
			dp[i] = min(dp[i], dp[max(i-lcs[i][j], j)]+y)
		}
	}
	Fprint(out, dp[n])
}

//func main() { CF1120C(os.Stdin, os.Stdout) }
