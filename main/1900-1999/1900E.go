package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"slices"
)

func CF1900E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for ; m > 0; m-- {
			var v, w int
			Fscan(in, &v, &w)
			g[v-1] = append(g[v-1], w-1)
		}

		allScc := [][]int{}
		dfn := make([]int, n)
		ts := 0
		st := []int{}
		var tarjan func(int) int
		tarjan = func(v int) int {
			ts++
			dfn[v] = ts
			lowV := ts
			st = append(st, v)
			for _, w := range g[v] {
				if dfn[w] == 0 {
					lowV = min(lowV, tarjan(w))
				} else {
					lowV = min(lowV, dfn[w])
				}
			}
			if dfn[v] == lowV {
				scc := []int{}
				for {
					w := st[len(st)-1]
					st = st[:len(st)-1]
					dfn[w] = math.MaxInt
					scc = append(scc, w)
					if w == v {
						break
					}
				}
				allScc = append(allScc, scc)
			}
			return lowV
		}
		for i, t := range dfn {
			if t == 0 {
				tarjan(i)
			}
		}
		slices.Reverse(allScc)

		ns := len(allScc)
		a2 := make([]int, ns)
		sid := make([]int, n)
		for i, cc := range allScc {
			for _, v := range cc {
				a2[i] += a[v]
				sid[v] = i
			}
		}

		g2 := make([][]int, ns)
		deg := make([]int, ns)
		for v, ws := range g {
			v = sid[v]
			for _, w := range ws {
				w = sid[w]
				if v != w {
					g2[v] = append(g2[v], w)
					deg[w]++
				}
			}
		}

		type pair struct{ len, s int }
		f := make([]pair, ns)
		q := make([]int, 0, ns)
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		ans := pair{}
		gr := func(p, q pair) bool { return p.len > q.len || p.len == q.len && p.s < q.s }
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			p := f[v]
			p.len += len(allScc[v])
			p.s += a2[v]
			if gr(p, ans) {
				ans = p
			}
			for _, w := range g2[v] {
				if gr(p, f[w]) {
					f[w] = p
				}
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		Fprintln(out, ans.len, ans.s)
	}
}

//func main() { CF1900E(os.Stdin, os.Stdout) }
