package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1446B(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var n, m, ans int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &m, &s, &t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i, x := range s {
		for j, y := range t {
			if x == y {
				dp[i+1][j+1] = dp[i][j] + 2
				ans = max(ans, dp[i+1][j+1])
			} else {
				dp[i+1][j+1] = max(max(dp[i][j+1], dp[i+1][j])-1, 0)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1446B(os.Stdin, os.Stdout) }
