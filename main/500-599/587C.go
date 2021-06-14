package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF587C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	k := 10
	merge := func(a, b []int) []int {
		i, n := 0, len(a)
		j, m := 0, len(b)
		c := make([]int, 0, n+m)
		for {
			if i == n {
				c = append(c, b[j:]...)
				break
			}
			if j == m {
				c = append(c, a[i:]...)
				break
			}
			if a[i] < b[j] {
				c = append(c, a[i])
				i++
			} else {
				c = append(c, b[j])
				j++
			}
		}
		if len(c) > k {
			c = c[:k]
		}
		return c
	}

	var n, m, q, v, w int
	Fscan(in, &n, &m, &q)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	const mx = 17
	type pair struct {
		p    int
		mins []int
	}
	pa := make([][mx]pair, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v)
		if v--; len(pa[v][0].mins) < 10 {
			pa[v][0].mins = append(pa[v][0].mins, i)
		}
	}
	dep := make([]int, n)
	var f func(v, p, d int)
	f = func(v, p, d int) {
		pa[v][0].p = p
		dep[v] = d
		for _, w := range g[v] {
			if w != p {
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 0)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1] = pair{pp.p, merge(p.mins, pp.mins)}
			} else {
				pa[v][i+1] = pair{p: -1}
			}
		}
	}

	for ; q > 0; q-- {
		Fscan(in, &v, &w, &k)
		v--
		w--
		if dep[v] > dep[w] {
			v, w = w, v
		}
		mins := []int{}
		for i := 0; i < mx; i++ {
			if (dep[w]-dep[v])>>i&1 > 0 {
				p := pa[w][i]
				mins = merge(mins, p.mins)
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					mins = merge(mins, merge(pv.mins, pw.mins))
					v, w = pv.p, pw.p
				}
			}
			mins = merge(mins, merge(pa[v][0].mins, pa[w][0].mins))
			v = pa[v][0].p
		}
		mins = merge(mins, pa[v][0].mins)
		Fprint(out, len(mins))
		for _, v := range mins {
			Fprint(out, " ", v)
		}
		Fprintln(out)
	}
}

//func main() { CF587C(os.Stdin, os.Stdout) }
