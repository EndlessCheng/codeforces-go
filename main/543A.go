package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF543A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, b, mod, v, s int
	Fscan(in, &n, &m, &b, &mod)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, b+1)
	}
	dp[0][0] = 1
	for ; n > 0; n-- {
		Fscan(in, &v)
		for j := 1; j <= m; j++ {
			for k := v; k <= b; k++ {
				dp[j][k] = (dp[j][k] + dp[j-1][k-v]) % mod
			}
		}
	}
	for _, v := range dp[m] {
		s = (s + v) % mod
	}
	Fprint(out, s)
}

//func main() { CF543A(os.Stdin, os.Stdout) }
