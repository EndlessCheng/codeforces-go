package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p3800(in io.Reader, out io.Writer) {
	var n, m, k, t, x, y, v int
	Fscan(in, &n, &m, &k, &t)
	type pair struct{ j, v int }
	a := make([][]pair, n)
	for ; k > 0; k-- {
		Fscan(in, &x, &y, &v)
		a[x-1] = append(a[x-1], pair{y - 1, v})
	}

	f := make([]int, m)
	for _, ps := range a {
		slices.SortFunc(ps, func(a, b pair) int { return a.j - b.j })
		nf := make([]int, m)
		q := []int{}
		for j, fv := range f[:min(t, m)] {
			for len(q) > 0 && fv >= f[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
		}
		for j := range nf {
			if j+t < m {
				for len(q) > 0 && f[j+t] >= f[q[len(q)-1]] {
					q = q[:len(q)-1]
				}
				q = append(q, j+t)
			}
			if q[0] < j-t {
				q = q[1:]
			}
			nf[j] = f[q[0]]
			if len(ps) > 0 && ps[0].j == j {
				nf[j] += ps[0].v
				ps = ps[1:]
			}
		}
		f = nf
	}
	Fprint(out, slices.Max(f))
}

//func main() { p3800(bufio.NewReader(os.Stdin), os.Stdout) }
