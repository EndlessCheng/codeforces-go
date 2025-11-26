package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func p10949(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type tuple struct{ v, w, wt int }
	es := make([]tuple, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
	}
	slices.SortFunc(es, func(a, b tuple) int { return a.wt - b.wt })

	m = 1 << n
	sum := make([]int, m)
	fa := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for mask := range sum {
		s := 0
		for m := uint(mask); m > 0; m &= m - 1 {
			s += a[bits.TrailingZeros(m)]
		}
		if s != 0 {
			sum[mask] = 1e9
			continue
		}

		for i := range fa {
			fa[i] = i
		}
		for _, e := range es {
			v, w := e.v, e.w
			if mask>>v&1 == 0 || mask>>w&1 == 0 {
				continue
			}
			fv, fw := find(v), find(w)
			if fv != fw {
				fa[fv] = fw
				sum[mask] += e.wt
			}
		}
	}
	f := make([]int, m)
	for i := 1; i < m; i++ {
		f[i] = 1e9
	}
	for s, fs := range f {
		t := m - 1 ^ s
		for sub := t; sub > 0; sub = (sub - 1) & t {
			ss := s | sub
			f[ss] = min(f[ss], fs+sum[sub])
		}
	}
	if f[m-1] == 1e9 {
		Fprint(out, "Impossible")
	} else {
		Fprint(out, f[m-1])
	}
}

//func main() { p10949(bufio.NewReader(os.Stdin), os.Stdout) }
