package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type xorBasis02 struct{ b, d [20]int32 }

func (b *xorBasis02) insert(v, d int32) bool {
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

func (b *xorBasis02) decompose(v, d int32) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i == 0 {
			continue
		}
		if b.b[i] == 0 || b.d[i] < d {
			return false
		}
		v ^= b.b[i]
	}
	return true
}

func (b *xorBasis02) merge(other *xorBasis02) {
	for i := len(other.b) - 1; i >= 0; i-- {
		x := other.b[i]
		if x > 0 {
			b.insert(x, other.d[i])
		}
	}
}

func cf1902F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, x, y int
	var k int32
	Fscan(in, &n)
	a := make([]int32, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mx = 18
	pa := make([][mx]int, n)
	dep := make([]int32, n)
	xb := make([]xorBasis02, n)
	var dfs func(int, int)
	dfs = func(v, p int) {
		pa[v][0] = p
		for _, w := range g[v] {
			if w == p {
				continue
			}
			dep[w] = dep[v] + 1
			xb[w] = xb[v]
			xb[w].insert(a[w], dep[w])
			dfs(w, v)
		}
	}
	xb[0].insert(a[0], 0)
	dfs(0, -1)
	for i := range mx - 1 {
		for v := range pa {
			p := pa[v][i]
			if p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v int, d int32) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros32(k)]
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

	Fscan(in, &q)
	for range q {
		Fscan(in, &x, &y, &k)
		x--
		y--
		b := xb[x]
		b.merge(&xb[y])
		if b.decompose(k, dep[getLCA(x, y)]) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1902F(bufio.NewReader(os.Stdin), os.Stdout) }
