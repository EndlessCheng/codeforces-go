package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf840E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const B = 256
	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	g := make([][]int, n+1)
	for range n - 1 {
		var u, v int
		Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	f := make([][B]int, n+1)
	fa := make([]int, n+1)
	dep := make([]int, n+1)
	jump := make([]int, n+1)

	var dfs func(int)
	dfs = func(v int) {
		dep[v] = dep[fa[v]] + 1
		for _, w := range g[v] {
			if w != fa[v] {
				fa[w] = v
				dfs(w)
			}
		}
		if dep[v] < B {
			return
		}

		u := v
		for i := range B {
			x := (i^a[u])>>8 ^ 255
			f[v][x] = max(f[v][x], 255<<8|(i^a[u]))
			u = fa[u]
		}
		jump[v] = u
		for i := range 8 {
			for j := range B {
				f[v][j] = max(f[v][j], f[v][j^1<<i]-B<<i)
			}
		}
	}

	dfs(1)

	for range q {
		var u, v int
		Fscan(in, &u, &v)
		d := dep[v] - dep[u] + 1
		ans := 0
		for i := range d / B {
			ans = max(ans, f[v][i])
			v = jump[v]
		}
		for i := d / B * B; i < d; i++ {
			ans = max(ans, i^a[v])
			v = fa[v]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf840E(bufio.NewReader(os.Stdin), os.Stdout) }
