package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1209E2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, m)
		for i := range a {
			a[i] = make([]int, n+1)
		}
		for j := 1; j <= n; j++ {
			for i, r := range a {
				Fscan(in, &r[j])
				a[i][0] = max(a[i][0], r[j])
			}
		}
		sort.Slice(a, func(i, j int) bool { return a[i][0] > a[j][0] })
		if m > n {
			m = n
			a = a[:n]
		}

		mx := make([][]int, m)
		for i, r := range a {
			r = r[1:]
			mx[i] = make([]int, 1<<n)
			sum := make([]int, 1<<n)
			for p, v := range r {
				for s := 0; s < 1<<p; s++ {
					t, w := 1<<p|s, sum[s]+v
					sum[t] = w
					for range r {
						mx[i][t] = max(mx[i][t], w)
						t = t<<1&(1<<n-1) | t>>(n-1)
					}
				}
			}
		}

		dp := make([][]int, m)
		for i := range dp {
			dp[i] = make([]int, 1<<n)
		}
		dp[0] = mx[0]
		for i := 1; i < m; i++ {
			for s, v := range mx[i] {
				for ss := s; ss < 1<<n; ss = (ss + 1) | s {
					dp[i][ss] = max(dp[i][ss], dp[i-1][ss^s]+v)
				}
			}
		}
		Fprintln(out, dp[m-1][1<<n-1])
	}
}

//func main() { CF1209E2(os.Stdin, os.Stdout) }
