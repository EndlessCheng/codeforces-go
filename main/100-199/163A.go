package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF163A(_r io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var s, t string
	Fscan(_r, &s, &t)
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i, v := range s {
		for j, w := range t {
			if v == w {
				dp[i+1][j+1] = (dp[i+1][j] + dp[i][j] + 1) % mod
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}
	ans := 0
	for _, dv := range dp {
		ans = (ans + dv[m]) % mod
	}
	Fprint(out, ans)
}

//func main() { CF163A(os.Stdin, os.Stdout) }
