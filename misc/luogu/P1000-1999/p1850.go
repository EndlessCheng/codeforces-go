package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1850(in io.Reader, out io.Writer) {
	var n, m, V, E int
	Fscan(in, &n, &m, &V, &E)
	m = min(m, n)
	type pair struct {
		v [2]int
		p float64
	}
	a := make([]pair, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].v[0])
	}
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].v[1])
	}
	for i := 0; i < n; i++ {
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

	f := make([][2]float64, m+1)
	for i := 1; i <= n; i++ {
		// 从 n 出发倒着走
		cur1 := a[i].v[0]
		cur2 := a[i].v[1]
		pre1 := a[i-1].v[0]
		pre2 := a[i-1].v[1]
		curP := a[i].p
		preP := a[i-1].p
		for j := m; j >= max(m-(n-i), 0); j-- {
			f[j][1] = f[j][0] + float64(g[cur1][pre1])*(1-curP) + float64(g[cur2][pre1])*curP
			if j > 0 {
				f[j][1] = min(f[j][1], f[j-1][1]+float64(g[cur1][pre1])*(1-curP)*(1-preP)+float64(g[cur1][pre2])*(1-curP)*preP+
					float64(g[cur2][pre1])*curP*(1-preP)+float64(g[cur2][pre2])*curP*preP)
			}
			f[j][0] += float64(g[cur1][pre1])
			if j > 0 {
				f[j][0] = min(f[j][0], f[j-1][1]+float64(g[cur1][pre1])*(1-preP)+float64(g[cur1][pre2])*preP)
			}
		}
	}
	Fprintf(out, "%.2f", f[m][0])
}

//func main() { p1850(bufio.NewReader(os.Stdin), os.Stdout) }
