package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	dp := make([][][]int, n+m)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	dp[0][0][0] = a[0][0]
	for s := 0; s < n+m-1; s++ {
		for r1 := 0; r1 < n && r1 <= s; r1++ {
			for r2 := 0; r2 < n && r2 <= s; r2++ {
				if r1 == 0 && r2 == 0 {
					continue
				}
				dv := &dp[s+1][r1+1][r2+1]
				*dv = max(max(max(dp[s][r1+1][r2+1], dp[s][r1+1][r2]), dp[s][r1][r2+1]), dp[s][r1][r2])
				if r1 == r2 {
					if s-r1 < m {
						*dv += a[r1][s-r1]
					}
				} else {
					if s-r1 < m {
						*dv += a[r1][s-r1]
					}
					if s-r2 < m {
						*dv += a[r2][s-r2]
					}
				}
			}
		}
	}
	Fprint(out, dp[n+m-1][n][n])
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
