package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF698B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type edge struct{ to, i int }

	var n, w int
	Fscan(in, &n)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) { fa[find(x)] = find(y) }
	same := func(x, y int) bool { return find(x) == find(y) }
	a := make([]interface{}, n+1)
	g := make([][]edge, n+1)
	hasRoot := false
	for i := 1; i <= n; i++ {
		Fscan(in, &w)
		if i == w {
			hasRoot = true
			g[i] = append(g[i], edge{i, i})
		} else {
			merge(i, w)
			g[i] = append(g[i], edge{w, i})
			g[w] = append(g[w], edge{i, i})
		}
		a[i] = w
	}

	vs := []int{}
	for i := 1; i <= n; i++ {
		if i == fa[find(i)] {
			vs = append(vs, i)
		}
	}
	ans := len(vs)
	if hasRoot {
		ans--
	}
	Fprintln(out, ans)

	skipOne := true
	vis := make([]int8, n+1)
	var f func(int)
	f = func(v int) {
		vis[v] = 1
		for _, e := range g[v] {
			w := e.to
			if t := vis[w]; t == 0 {
				f(w)
			} else if t == 2 || w == v {
				// 无根造根，有根跳根
				if t == 2 && !hasRoot {
					hasRoot = true
					a[e.i] = e.i
					continue
				}
				if w == v && skipOne {
					skipOne = false
					continue
				}
				// 一个合并 CC 的技巧
				if same(e.i, vs[0]) {
					vs[0], vs[1] = vs[1], vs[0]
				}
				merge(e.i, vs[0])
				a[e.i], vs = vs[0], vs[1:]
			}
		}
		vis[v] = 2
	}
	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			f(i)
		}
	}
	Fprint(out, a[1:]...)
}

//func main() { CF698B(os.Stdin, os.Stdout) }
