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
	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	g := make([][]int, n)
	for range n - 1 {
		var v, u int
		Fscan(in, &v, &u)
		v--
		u--
		g[v] = append(g[v], u)
		g[u] = append(g[u], v)
	}

	type pair struct{ x, y int }
	qs := make([][]pair, n)
	for i := range q {
		var u, v int
		Fscan(in, &u, &v)
		u--
		v--
		qs[v] = append(qs[v], pair{u, i})
	}

	ans := make([]int, q)
	dep := make([]int, n)
	st := make([]int, n)
	top := n

	var dfs func(int, int, int)
	dfs = func(v, fa, d int) {
		dep[v] = d
		top--
		st[top] = a[v]

		for _, qu := range qs[v] {
			res := 0
			for i := range dep[v] - dep[qu.x] + 1 {
				res = max(res, i^st[top+i])
			}
			ans[qu.y] = res
		}

		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
			}
		}

		top++
	}

	dfs(0, -1, 0)

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf840E(bufio.NewReader(os.Stdin), os.Stdout) }
