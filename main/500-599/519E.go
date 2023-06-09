package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF519E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, q int
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
	size := make([]int, n)
	var f func(int, int) int
	f = func(v, p int) int {
		pa[v][0] = p
		sz := 1
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				sz += f(w, v)
			}
		}
		size[v] = sz
		return sz
	}
	f(0, -1)
	for k := 0; k+1 < mx; k++ {
		for v := range pa {
			if p := pa[v][k]; p != -1 {
				pa[v][k+1] = pa[p][k]
			} else {
				pa[v][k+1] = -1
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
		if v == w {
			return v
		}
		for k := mx - 1; k >= 0; k-- {
			if pa[v][k] != pa[w][k] {
				v, w = pa[v][k], pa[w][k]
			}
		}
		return pa[v][0]
	}

	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &w)
		if v == w {
			Fprintln(out, n)
			continue
		}
		v--
		w--
		if dep[v] > dep[w] {
			v, w = w, v
		}
		lca := _lca(v, w)
		d := dep[v] + dep[w] - dep[lca]<<1
		if d&1 == 1 {
			Fprintln(out, 0)
			continue
		}
		uw := uptoDep(w, dep[w]-d>>1+1)
		if p := pa[uw][0]; p != lca {
			Fprintln(out, size[p]-size[uw])
		} else {
			uv := uptoDep(v, dep[v]-d>>1+1)
			Fprintln(out, n-size[uv]-size[uw])
		}
	}
}

//func main() { CF519E(os.Stdin, os.Stdout) }
