package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1517D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type edge struct{ x, y, wt int }

	var n, m, k, wt int
	Fscan(in, &n, &m, &k)
	if k&1 > 0 {
		for s := strings.Repeat("-1 ", m); n > 0; n-- {
			Fprintln(out, s)
		}
		return
	}
	k /= 2
	g := make([][][]edge, n)
	for i := range g {
		g[i] = make([][]edge, m)
		for j := 1; j < m; j++ {
			Fscan(in, &wt)
			g[i][j-1] = append(g[i][j-1], edge{i, j, wt})
			g[i][j] = append(g[i][j], edge{i, j - 1, wt})
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &wt)
			g[i-1][j] = append(g[i-1][j], edge{i, j, wt})
			g[i][j] = append(g[i][j], edge{i - 1, j, wt})
		}
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(x, y, k int) int
	f = func(x, y, k int) (res int) {
		if k == 0 {
			return
		}
		dv := &dp[x][y][k]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for _, e := range g[x][y] {
			if r := f(e.x, e.y, k-1) + e.wt; r < res {
				res = r
			}
		}
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fprint(out, f(i, j, k)*2, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1517D(os.Stdin, os.Stdout) }
