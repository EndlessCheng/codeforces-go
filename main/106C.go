package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF106C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxW, n, a, b int
	Fscan(in, &maxW, &n)
	n++
	weights := make([]int, n)
	values := make([]int, n)
	stocks := make([]int, n)
	Fscan(in, &weights[0], &values[0])
	stocks[0] = maxW / weights[0]
	for i := 1; i < n; i++ {
		Fscan(in, &a, &b, &weights[i], &values[i])
		stocks[i] = a / b
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i, vi := range values {
		si, wi := stocks[i], weights[i]
		for j := range dp[i] {
			for k := 0; k <= si && k*wi <= j; k++ {
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-k*wi]+k*vi)
			}
		}
	}
	Fprint(_w, dp[n][maxW])
}

//func main() { CF106C(os.Stdin, os.Stdout) }
