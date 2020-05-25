package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 原题 https://www.luogu.com.cn/problem/P1854
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	from := make([][]int, n)
	for i := range from {
		from[i] = make([]int, m)
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1e18
		}
	}
	dp[0][0] = a[0][0]
	for j := 1; j < m; j++ {
		dp[0][j] = max(dp[0][j-1], a[0][j])
	}
	for i := 1; i < n; i++ {
		for j := i; j < m; j++ {
			mx := int(-1e18)
			for k := j; k >= i; k-- {
				mx = max(mx, a[i][k])
				if dp[i-1][k-1]+mx > dp[i][j] {
					dp[i][j] = dp[i-1][k-1] + mx
					from[i][j] = k - 1
				}
			}
		}
	}
	Fprintln(out, dp[n-1][m-1])
	ans := make([]int, n)
	end := m - 1
	for i := n - 1; i > 0; i-- {
		st := from[i][end]
		ans[i] = end
		for j := st + 1; j < end; j++ {
			if a[i][j] > a[i][ans[i]] {
				ans[i] = j
			}
		}
		end = st
	}
	for j := 1; j <= end; j++ {
		if a[0][j] > a[0][ans[0]] {
			ans[0] = j
		}
	}
	for i, v := range ans {
		if i > 0 {
			Fprint(out, " ")
		}
		Fprint(out, v+1)
	}
}

func main() { run(os.Stdin, os.Stdout) }
