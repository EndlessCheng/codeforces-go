package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF455C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, v, w, u, core, maxD, op int
	Fscan(in, &n, &m, &q)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	size := make([]int, n)
	fa := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		sf, st := size[from], size[to]
		s1, s2 := (sf+1)/2, (st+1)/2
		if s1 > s2 {
			s1, s2 = s2, s1
		}
		if s1+1 >= s2 {
			size[to] = s1 + s2 + 1
		} else if sf > st {
			size[to] = sf
		} else {
			size[to] = st
		}
		fa[from] = to
	}

	vis := make([]bool, n)
	var f func(v, p, d int)
	f = func(v, p, d int) {
		vis[v] = true
		fa[v] = core
		if d > maxD {
			maxD = d
			u = v
		}
		for _, w := range g[v] {
			if w != p {
				f(w, v, d+1)
			}
		}
	}
	for i, b := range vis {
		if !b {
			core = i
			maxD = -1
			f(i, -1, 0)
			maxD = -1
			f(u, -1, 0)
			size[i] = maxD
		}
	}

	for ; q > 0; q-- {
		Fscan(in, &op, &v)
		if op == 1 {
			Fprintln(out, size[find(v-1)])
		} else {
			Fscan(in, &w)
			merge(v-1, w-1)
		}
	}
}

//func main() { CF455C(os.Stdin, os.Stdout) }
