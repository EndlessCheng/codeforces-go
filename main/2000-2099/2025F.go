package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2025F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, m+1)
	b := make([]int, m+1)
	c := make([]int, n+1)
	type edge struct{ to, i int }
	g := make([][]edge, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], edge{b[i], i})
		g[b[i]] = append(g[b[i]], edge{a[i], i})
		c[a[i]] ^= 1
	}

	ans := make([]bool, m+1)
	vis := make([]bool, n+1)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, e := range g[v] {
			w := e.to
			if vis[w] {
				continue
			}
			dfs(w)
			if c[w] != 0 {
				ans[e.i] = true
				c[v] ^= 1
				c[w] ^= 1
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			dfs(i)
			c[i] = 0
		}
	}

	for i := 1; i <= m; i++ {
		var x int
		if !ans[i] {
			Fprint(out, "x")
			x = a[i]
		} else {
			Fprint(out, "y")
			x = b[i]
		}
		if c[x] != 0 {
			Fprintln(out, "-")
		} else {
			Fprintln(out, "+")
		}
		c[x] ^= 1
	}
}

//func main() { cf2025F(bufio.NewReader(os.Stdin), os.Stdout) }
