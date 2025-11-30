package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf842E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, p int
	Fscan(in, &n)
	n++

	const mx = 19
	pa := make([][mx]int, n)
	dep := make([]int, n)
	uptoDep := func(v, d int) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros32(k)]
		}
		return v
	}
	lca := func(v, w int) int {
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
	dis := func(v, w int) int { return dep[v] + dep[w] - dep[lca(v, w)]*2 }

	a, b := []int{0}, []int{1}
	dia := 1
	for i := 1; i < n; i++ {
		Fscan(in, &p)
		p--
		dep[i] = dep[p] + 1
		pa[i][0] = p
		for j := range mx - 1 {
			pa[i][j+1] = pa[pa[i][j]][j]
		}
		if i > 1 {
			da, db := dis(i, a[0]), dis(i, b[0])
			if da < db {
				da, db = db, da
				a, b = b, a
			}
			if da > dia {
				dia = da
				for _, v := range b {
					if dis(i, v) == dia {
						a = append(a, v)
					}
				}
				b = []int{i}
			} else if da == dia {
				b = append(b, i)
			}
		}
		Fprintln(out, len(a)+len(b))
	}
}

//func main() { cf842E(bufio.NewReader(os.Stdin), os.Stdout) }
