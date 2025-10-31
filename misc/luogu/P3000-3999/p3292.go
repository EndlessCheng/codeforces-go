package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type xorBasis292 struct {
	b [61]int
	d [61]int16
}

func (b *xorBasis292) insert(v int, d int16) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.d[i] = d
			return true
		}
		if d > b.d[i] {
			d, b.d[i] = b.d[i], d
			v, b.b[i] = b.b[i], v
		}
		v ^= b.b[i]
	}
	return false
}

func (b *xorBasis292) maxXor(lower int16) (res int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if res>>i&1 == 0 && b.d[i] >= lower {
			res = max(res, res^b.b[i])
		}
	}
	return
}

func (b *xorBasis292) merge(other *xorBasis292) {
	for i := len(other.b) - 1; i >= 0; i-- {
		x := other.b[i]
		if x > 0 {
			b.insert(x, other.d[i])
		}
	}
}

func p3292(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, x, y int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mx = 15
	pa := make([][mx]int, n)
	dep := make([]int16, n)
	xb := make([]xorBasis292, n)
	var dfs func(int, int)
	dfs = func(v, p int) {
		xb[v].insert(a[v], dep[v])
		pa[v][0] = p
		for _, w := range g[v] {
			if w == p {
				continue
			}
			dep[w] = dep[v] + 1
			xb[w] = xb[v]
			dfs(w, v)
		}
	}
	dfs(0, -1)
	for i := 0; i < mx-1; i++ {
		for v := range pa {
			p := pa[v][i]
			if p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v int, d int16) int {
		for k := uint16(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros16(k)]
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
			pv, pw := pa[v][i], pa[w][i]
			if pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	for ; q > 0; q-- {
		Fscan(in, &x, &y)
		x--
		y--
		b := xb[x]
		b.merge(&xb[y])
		Fprintln(out, b.maxXor(dep[getLCA(x, y)]))
	}
}

//func main() { p3292(bufio.NewReader(os.Stdin), os.Stdout) }
