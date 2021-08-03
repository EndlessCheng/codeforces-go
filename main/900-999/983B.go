package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF983B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, q, l, r int
	Fscan(in, &n)
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
		Fscan(in, &dp[i][i])
	}
	for i := n - 1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1] ^ dp[i+1][j]
		}
	}
	for i := n - 1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = max(dp[i][j], max(dp[i][j-1], dp[i+1][j]))
		}
	}
	Fscan(in, &q)
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, dp[l][r])
	}
}

//func main() { CF983B(os.Stdin, os.Stdout) }
