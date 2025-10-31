package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func p1879(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9

	var n, m, v int
	Fscan(in, &n, &m)
	g := make([]int, n)
	for i := range g {
		for j := m - 1; j >= 0; j-- {
			Fscan(in, &v)
			g[i] |= v << uint(j)
		}
	}
	dp := make([][]int, n)
	dp[0] = make([]int, 1<<uint(m))
	for j := range dp[0] {
		if j<<1&j == 0 && j|g[0] == g[0] {
			dp[0][j] = 1
		}
	}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 1<<uint(m))
		for j := range dp[i] {
			if j<<1&j == 0 && j|g[i] == g[i] {
				for k, v := range dp[i-1] {
					if j&k == 0 {
						dp[i][j] += v
					}
				}
			}
			dp[i][j] %= mod
		}
	}
	sum := 0
	for _, v := range dp[n-1] {
		sum += v
	}
	Fprint(out, sum%mod)
}

//func main() { p1879(os.Stdin, os.Stdout) }
