package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func p1850(in io.Reader, out io.Writer) {
	var n, m, V, E int
	Fscan(in, &n, &m, &V, &E)
	type pair struct {
		v [2]int
		p float64
	}
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].v[0])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].v[1])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].p)
	}

	g := make([][]int, V+1)
	for i := range g {
		g[i] = make([]int, V+1)
		if i == 0 {
			continue
		}
		for j := range g[i] {
			if j != i {
				g[i][j] = 1e18
			}
		}
	}
	for ; E > 0; E-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		g[v][w] = min(g[v][w], wt)
		g[w][v] = min(g[w][v], wt)
	}
	for k := range g {
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}

	dp := make([][][2]float64, n)
	for i := range dp {
		dp[i] = make([][2]float64, m+1)
		for j := range dp[i] {
			dp[i][j] = [2]float64{-1, -1}
		}
	}
	var f func(int, int, int) float64
	f = func(i, j, k int) (res float64) {
		if i == n {
			return 0
		}
		p := &dp[i][j][k]
		if *p != -1 {
			return *p
		}
		curP := a[i].p
		nxtP := a[i+1].p
		nxt1 := a[i+1].v[0]
		nxt2 := a[i+1].v[1]
		if k == 0 {
			cur := a[i].v[0]
			res = f(i+1, j, 0) + float64(g[cur][nxt1])
			if j > 0 {
				res = min(res, f(i+1, j-1, 1)+float64(g[cur][nxt1])*(1-nxtP)+float64(g[cur][nxt2])*nxtP)
			}
		} else {
			cur1 := a[i].v[0]
			cur2 := a[i].v[1]
			res = f(i+1, j, 0) + float64(g[cur1][nxt1])*(1-curP) + float64(g[cur2][nxt1])*curP
			if j > 0 {
				res = min(res, f(i+1, j-1, 1)+float64(g[cur1][nxt1])*(1-curP)*(1-nxtP)+float64(g[cur1][nxt2])*(1-curP)*nxtP+
					float64(g[cur2][nxt1])*curP*(1-nxtP)+float64(g[cur2][nxt2])*curP*nxtP)
			}
		}
		*p = res
		return
	}
	Fprintf(out, "%.2f", f(0, m, 0))
}

func main() { p1850(bufio.NewReader(os.Stdin), os.Stdout) }
