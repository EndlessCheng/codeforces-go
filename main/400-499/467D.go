package main

import (
	. "fmt"
	"io"
	"math"
	"slices"
	"strings"
)

// https://github.com/EndlessCheng
func cf467D(in io.Reader, out io.Writer) {
	var n, m, ts, ansR, ansSz int
	var s, t string
	Fscan(in, &n)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &m)
	idx := map[string]int{}
	type pair struct{ r, sz int }
	less := func(a, b pair) bool { return a.r < b.r || a.r == b.r && a.sz < b.sz }
	data := make([]pair, m*2)
	get := func(s string) int {
		s = strings.ToLower(s)
		if _, ok := idx[s]; !ok {
			data[len(idx)] = pair{strings.Count(s, "r"), len(s)}
			idx[s] = len(idx)
		}
		return idx[s]
	}
	g := make([][]int, m*2)
	for range m {
		Fscan(in, &s, &t)
		v, w := get(s), get(t)
		g[v] = append(g[v], w)
	}
	g = g[:len(idx)]

	allScc := [][]int{}
	dfn := make([]int, len(g))
	st := []int{}
	var tarjan func(int) int
	tarjan = func(v int) int {
		ts++
		dfn[v] = ts
		lowV := ts
		st = append(st, v)
		for _, w := range g[v] {
			if dfn[w] == 0 {
				lowW := tarjan(w)
				lowV = min(lowV, lowW)
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

	sid := make([]int, len(g))
	f := make([]pair, len(allScc))
	for i, scc := range allScc {
		mn := pair{1e9, 0}
		for _, v := range scc {
			sid[v] = i
			p := data[v]
			if less(p, mn) {
				mn = p
			}
		}
		f[i] = mn
	}
	g2 := make([][]int, len(allScc))
	deg := make([]int, len(allScc))
	for v, ws := range g {
		v = sid[v]
		for _, w := range ws {
			w = sid[w]
			if v != w {
				g2[w] = append(g2[w], v)
				deg[v]++
			}
		}
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g2[v] {
			if less(f[v], f[w]) {
				f[w] = f[v]
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	for _, s := range a {
		s = strings.ToLower(s)
		if v, ok := idx[s]; ok {
			p := f[sid[v]]
			ansR += p.r
			ansSz += p.sz
		} else {
			ansR += strings.Count(s, "r")
			ansSz += len(s)
		}
	}
	Fprint(out, ansR, ansSz)
}

//func main() { cf467D(bufio.NewReader(os.Stdin), os.Stdout) }
