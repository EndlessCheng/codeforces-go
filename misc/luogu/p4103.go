package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func p4103(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n := r()
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		v, w := r()-1, r()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mx = 20
	pa := make([][mx]int, n)
	dep := make([]int, n)
	dfn := make([]int, n)
	ts := 0
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
	inNodes := make([]int, n)
	addVtEdge := func(v, w int) { vt[v] = append(vt[v], w) }
	st := []int{0}
	for q := r(); q > 0; q-- {
		k := r()
		nodes := make([]int, k)
		for i := range nodes {
			nodes[i] = r() - 1
		}
		sort.Slice(nodes, func(i, j int) bool { return dfn[nodes[i]] < dfn[nodes[j]] })
		vt[0] = vt[0][:0]
		st = st[:1]
		for _, v := range nodes {
			inNodes[v] = q
			if v == 0 {
				continue
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
			addVtEdge(st[i-1], st[i])
		}

		sum, gMinL, gMaxL := 0, int(1e9), 0
		var f func(int) (int, int, int)
		f = func(v int) (size, minL, maxL int) {
			imp := inNodes[v] == q
			if imp {
				size = 1
			}
			minL = 1e9
			for _, w := range vt[v] {
				sz, mn, mx := f(w)
				wt := dep[w] - dep[v]
				sum += wt * sz * (k - sz)
				size += sz
				mx += wt
				gMaxL = max(gMaxL, maxL+mx)
				maxL = max(maxL, mx)
				mn += wt
				if imp {
					gMinL = min(gMinL, mn)
				} else {
					gMinL = min(gMinL, minL+mn)
					minL = min(minL, mn)
				}
			}
			if minL == 1e9 {
				minL = 0
			}
			return
		}
		rt := 0
		if inNodes[0] != q && len(vt[0]) == 1 {
			rt = vt[0][0]
		}
		f(rt)
		Fprintln(out, sum, gMinL, gMaxL)
	}
}

//func main() { p4103(os.Stdin, os.Stdout) }
