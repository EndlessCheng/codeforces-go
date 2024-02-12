package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF2B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	zx := -1
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			Fscan(in, &v)
			if v == 0 {
				zx = i
			}
			for k, d := range []int{2, 5} {
				if i == 0 && j == 0 {
				} else if i == 0 {
					dp[i][j][k] = dp[i][j-1][k]
				} else if j == 0 {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = min(dp[i-1][j][k], dp[i][j-1][k])
				}
				if v > 0 {
					for ; v%d == 0; v /= d {
						dp[i][j][k]++
					}
				}
			}
		}
	}

	d := dp[n-1][n-1]
	k := 0
	if d[1] < d[0] {
		k = 1
	}
	if d[k] > 1 && zx != -1 {
		Fprint(out, "1\n", strings.Repeat("D", zx), strings.Repeat("R", n-1), strings.Repeat("D", n-1-zx))
		return
	}
	Fprintln(out, d[k])
	var f func(x, y int)
	f = func(x, y int) {
		if x == 0 && y == 0 {
			return
		}
		if x == 0 || y > 0 && dp[x][y-1][k] < dp[x-1][y][k] {
			f(x, y-1)
			Fprint(out, "R")
		} else {
			f(x-1, y)
			Fprint(out, "D")
		}
	}
	f(n-1, n-1)
}

//func main() { CF2B(os.Stdin, os.Stdout) }
