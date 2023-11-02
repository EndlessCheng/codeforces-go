package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func CF613D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, ts, q, k int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	const mx = 17
	dfn := make([]int, n)
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var buildPa func(int, int)
	buildPa = func(v, p int) {
		dfn[v] = ts
		ts++
		pa[v][0] = p
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				buildPa(w, v)
			}
		}
	}
	buildPa(0, -1)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := dep[v] - d; k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(uint(k))]
		}
		return v
	}
	_lca := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	g2 := make([][]int, n)
	st := []int{0}
	imp := make([]int, n)
o:
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &k)
		vs := make([]int, k)
		for i := range vs {
			Fscan(in, &vs[i])
			vs[i]--
		}
		sort.Slice(vs, func(i, j int) bool { return dfn[vs[i]] < dfn[vs[j]] })
		g2[0] = g2[0][:0]
		st = st[:1]
		for _, v := range vs {
			imp[v] = q
			if v == 0 {
				continue
			}
			if imp[pa[v][0]] == q {
				Fprintln(out, -1)
				continue o
			}
			g2[v] = g2[v][:0]
			lca := _lca(st[len(st)-1], v)
			if lca != st[len(st)-1] {
				for dfn[st[len(st)-2]] > dfn[lca] {
					top := st[len(st)-1]
					st = st[:len(st)-1]
					p := st[len(st)-1]
					g2[p] = append(g2[p], top)
				}
				if lca != st[len(st)-2] {
					g2[lca] = g2[lca][:0]
					g2[lca] = append(g2[lca], st[len(st)-1])
					st[len(st)-1] = lca
				} else {
					g2[lca] = append(g2[lca], st[len(st)-1])
					st = st[:len(st)-1]
				}
			}
			st = append(st, v)
		}
		for i := 1; i < len(st); i++ {
			g2[st[i-1]] = append(g2[st[i-1]], st[i])
		}

		ans := 0
		var f func(int) int
		f = func(v int) int {
			res := 0
			for _, w := range g2[v] {
				res += f(w)
			}
			if imp[v] == q {
				ans += res
				return 1
			}
			if res > 1 {
				ans++
				return 0
			}
			return res
		}
		f(0)
		Fprintln(out, ans)
	}
}

//func main() { CF613D(os.Stdin, os.Stdout) }
