package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1515F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, x, v, w, s int
	Fscan(in, &n, &m, &x)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	if s < (n-1)*x {
		Fprint(out, "NO")
		return
	}
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], edge{w, i})
		g[w] = append(g[w], edge{v, i})
	}

	Fprintln(out, "YES")
	todo := []int{}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, e := range g[v] {
			w := e.to
			if vis[w] {
				continue
			}
			dfs(w)
			if a[w] >= x {
				Fprintln(out, e.eid)
				a[v] += a[w] - x
			} else {
				todo = append(todo, e.eid)
			}
		}
	}
	dfs(0)
	for i := len(todo) - 1; i >= 0; i-- {
		Fprintln(out, todo[i])
	}
}

//func main() { cf1515F(os.Stdin, os.Stdout) }
