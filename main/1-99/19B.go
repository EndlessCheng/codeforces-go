package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF19B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
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

	var n, t, c int
	Fscan(in, &n)
	dp := make([]int64, n+1)
	for i := range dp {
		dp[i] = 1e18
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		Fscan(in, &t, &c)
		t++
		for j := n; j >= 0; j-- {
			dp[j] = min(dp[j], dp[max(j-t, 0)]+int64(c))
		}
	}
	Fprint(out, dp[n])
}

//func main() { CF19B(os.Stdin, os.Stdout) }
