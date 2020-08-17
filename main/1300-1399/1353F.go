package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1353F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		a := make([][]int64, n)
		for i := range a {
			a[i] = make([]int64, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}
		ans := int64(1e18)
		for i, row := range a {
			for j, v := range row {
				x := v - int64(i+j)
				if x > a[0][0] {
					continue
				}
				dp := make([][]int64, n)
				for i := range dp {
					dp[i] = make([]int64, m)
					for j := range dp[i] {
						dp[i][j] = 1e18
					}
				}
				dp[0][0] = a[0][0] - x
				for i, row := range a {
					for j, v := range row {
						d := v - x - int64(i+j)
						if d < 0 {
							continue
						}
						if i == 0 {
							if j > 0 {
								dp[i][j] = dp[i][j-1] + d
							}
						} else {
							if j == 0 {
								dp[i][j] = dp[i-1][j] + d
							} else {
								dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + d
							}
						}
					}
				}
				ans = min(ans, dp[n-1][m-1])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1353F(os.Stdin, os.Stdout) }
