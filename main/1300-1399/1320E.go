package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func CF1320E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, ts, q, k, m int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mx = 18
	pa := make([][mx]int, n)
	dep := make([]int, n)
	dfn := make([]int, n)
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
	getLCA := func(v, w int) int {
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
	getDis := func(v, w int) int { return dep[v] + dep[w] - dep[getLCA(v, w)]*2 }

	ans := make([]int, n)
	vt := make([][]int, n)
	ord := make([]int, n)
	spd := make([]int, n)
	addVtEdge := func(v, w int) { vt[v] = append(vt[v], w) }
	st := []int{0}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &k, &m)
		nodes := make([]int, k, k+m)
		for i := range nodes {
			Fscan(in, &nodes[i])
			nodes[i]--
			Fscan(in, &spd[nodes[i]])
			ord[nodes[i]] = i + 1
		}
		qs := make([]int, m)
		for i := range qs {
			Fscan(in, &qs[i])
			qs[i]--
			if ord[qs[i]] == 0 {
				nodes = append(nodes, qs[i])
			}
		}

		sort.Slice(nodes, func(i, j int) bool { return dfn[nodes[i]] < dfn[nodes[j]] })
		vt[0] = vt[0][:0]
		st = st[:1]
		for _, v := range nodes {
			if v == 0 {
				continue
			}
			vt[v] = vt[v][:0]
			lca := getLCA(st[len(st)-1], v)
			for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
				addVtEdge(st[len(st)-2], st[len(st)-1])
				st = st[:len(st)-1]
			}
			if lca != st[len(st)-1] {
				vt[lca] = vt[lca][:0]
				addVtEdge(lca, st[len(st)-1])
				st[len(st)-1] = lca
			}
			st = append(st, v)
		}
		for i := 1; i < len(st); i++ {
			addVtEdge(st[i-1], st[i])
		}

		var f func(int)
		f = func(v int) {
			av := -1
			minD := int(1e9)
			if ord[v] > 0 {
				av = v
				minD = 0
			}
			for _, w := range vt[v] {
				f(w)
				aw := ans[w]
				if aw < 0 {
					continue
				}
				d := (getDis(v, aw) + spd[aw] - 1) / spd[aw]
				if d < minD || d == minD && ord[aw] < ord[av] {
					minD = d
					av = aw
				}
			}
			ans[v] = av
		}
		var reroot func(int)
		reroot = func(v int) {
			av := ans[v]
			for _, w := range vt[v] {
				aw := ans[w]
				if aw < 0 {
					ans[w] = av
				} else {
					d := (getDis(w, av) + spd[av] - 1) / spd[av]
					d2 := (getDis(w, aw) + spd[aw] - 1) / spd[aw]
					if d < d2 || d == d2 && ord[av] < ord[aw] {
						ans[w] = av
					}
				}
				reroot(w)
			}
		}
		f(0)
		reroot(0)
		for _, v := range qs {
			Fprint(out, ord[ans[v]], " ")
		}
		Fprintln(out)
		for _, v := range nodes {
			ord[v] = 0
		}
	}
}

//func main() { CF1320E(os.Stdin, os.Stdout) }
