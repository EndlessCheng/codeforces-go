package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1166F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, v, w, c int
	var op string
	Fscanln(in, &n, &m, &q, &q)

	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	g := make([]map[int]bool, n+1)
	c2w := make([]map[int]int, n+1)
	for i := range g {
		g[i] = map[int]bool{}
		c2w[i] = map[int]int{}
	}
	add := func(v, w int) {
		g[find(v)][w] = true
		if c2w[v][c] == 0 {
			c2w[v][c] = w
			return
		}
		v = find(c2w[v][c])
		w = find(w)
		if v == w {
			return
		}
		if len(g[v]) > len(g[w]) {
			v, w = w, v
		}
		for x := range g[v] {
			g[w][x] = true
		}
		fa[v] = w
	}

	for range m {
		Fscanln(in, &v, &w, &c)
		add(v, w)
		add(w, v)
	}
	for range q {
		Fscanln(in, &op, &v, &w, &c)
		if op == "+" {
			add(v, w)
			add(w, v)
		} else if find(v) == find(w) || g[find(v)][w] {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { cf1166F(bufio.NewReader(os.Stdin), os.Stdout) }
