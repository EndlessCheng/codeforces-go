package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p1499(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	const mx = 17
	pa := make([][mx]int, n)
	dep := make([]int, n)
	dis := make([]int, n)
	type tuple struct{ fi, se, w int }
	downDis := make([]tuple, n)
	var build func(int, int)
	build = func(v, p int) {
		pa[v][0] = p
		fi, se, fw := 0, 0, -1
		for _, e := range g[v] {
			w := e.to
			if w == p {
				continue
			}
			dep[w] = dep[v] + 1
			dis[w] = dis[v] + e.wt
			build(w, v)
			d := downDis[w].fi + e.wt
			if d > fi {
				se = fi
				fi, fw = d, w
			} else if d > se {
				se = d
			}
		}
		downDis[v] = tuple{fi, se, fw}
	}
	build(0, -1)
	for i := 0; i < mx-1; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}

	upDis := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(v, fa, maxD int) {
		upDis[v] = maxD
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			if w != downDis[v].w {
				dfs(w, v, max(maxD, downDis[v].fi)+e.wt)
			} else {
				dfs(w, v, max(maxD, downDis[v].se)+e.wt)
			}
		}
	}
	dfs(0, -1, 0)

	uptoDep := func(v, d int) int {
		for k := uint(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
		}
		return v
	}
	getLCA := func(v, w int) (int, bool) {
		w = uptoDep(w, dep[v])
		if w == v {
			return v, false
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0], true
	}
	getDis := func(v, w int) int {
		if v == w {
			return upDis[v] + downDis[v].fi
		}
		if dep[v] > dep[w] {
			v, w = w, v
		}
		lca, ok := getLCA(v, w)
		d := dis[v] + dis[w] - dis[lca]*2 + downDis[w].fi
		if ok {
			return d + downDis[v].fi
		}
		if uptoDep(w, dep[v]+1) != downDis[v].w {
			return d + max(upDis[v], downDis[v].fi)
		}
		return d + max(upDis[v], downDis[v].se)
	}
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		Fprintln(out, getDis(v-1, w-1))
	}
}

//func main() { p1499(bufio.NewReader(os.Stdin), os.Stdout) }
