package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF598E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := [31][31][51]int{}
	var f func(n, m, k int) int
	f = func(n, m, k int) (ans int) {
		if n*m == k {
			return 0
		}
		v := &dp[n][m][k]
		if *v > 0 {
			return *v
		}
		defer func() { *v = ans }()
		ans = 1e9
		for i := 1; i < m; i++ {
			if n*i >= k {
				ans = min(ans, f(n, i, k)+n*n)
			} else {
				ans = min(ans, f(n, m-i, k-n*i)+n*n)
			}
		}
		for i := 1; i < n; i++ {
			if m*i >= k {
				ans = min(ans, f(i, m, k)+m*m)
			} else {
				ans = min(ans, f(n-i, m, k-m*i)+m*m)
			}
		}
		return
	}
	var t, n, m, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m, &k)
		Fprintln(out, f(n, m, k))
	}
}

//func main() { CF598E(os.Stdin, os.Stdout) }
