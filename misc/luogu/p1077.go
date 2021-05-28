package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func p1077(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int = 1e6 + 7

	var n, m, c int
	Fscan(in, &n, &m)
	dp := make([]int, m+1)
	dp[0] = 1
	for ; n > 0; n-- {
		Fscan(in, &c)
		for j := m; j >= 0; j-- {
			// 注：可以用前缀和优化
			for k := 1; k <= c && k <= j; k++ {
				dp[j] = (dp[j] + dp[j-k]) % mod
			}
		}
	}
	Fprint(out, dp[m])
}

//func main() { p1077(os.Stdin, os.Stdout) }
