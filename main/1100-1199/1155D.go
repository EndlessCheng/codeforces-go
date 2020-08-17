package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1155D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	var n, x, v, s int64
	Fscan(in, &n, &x)
	dp := [3]int64{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		dp[0] = max(0, dp[0]+v)
		dp[1] = max(dp[0], dp[1]+v*x)
		dp[2] = max(dp[1], dp[2]+v)
		s = max(s, dp[2])
	}
	Fprint(out, s)
}

//func main() { CF1155D(os.Stdin, os.Stdout) }
