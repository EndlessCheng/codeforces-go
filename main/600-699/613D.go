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

	vt := make([][]int, n)
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
		vt[0] = vt[0][:0]
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
			vt[v] = vt[v][:0]
			lca := getLCA(st[len(st)-1], v)
			for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
				p := st[len(st)-2]
				vt[p] = append(vt[p], st[len(st)-1])
				st = st[:len(st)-1]
			}
			if lca != st[len(st)-1] {
				vt[lca] = vt[lca][:0]
				vt[lca] = append(vt[lca], st[len(st)-1])
				st[len(st)-1] = lca
			}
			st = append(st, v)
		}
		for i := 1; i < len(st); i++ {
			vt[st[i-1]] = append(vt[st[i-1]], st[i])
		}

		ans := 0
		var f func(int) int
		f = func(v int) int {
			res := 0
			for _, w := range vt[v] {
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
