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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	us := make([]int, q+1)
	vs := make([]int, q+1)
	qid := make([][]int, n+1)
	for i := 1; i <= q; i++ {
		Fscan(in, &us[i], &vs[i])
		qid[vs[i]] = append(qid[vs[i]], i)
	}

	dep := make([]int, n+1)
	val := make([]int, n+1)
	ans := make([]int, q+1)

	var dfs func(int, int)
	dfs = func(v, fa int) {
		val[dep[v]] = a[v]
		for _, i := range qid[v] {
			res := 0
			for j := dep[us[i]]; j <= dep[v]; j++ {
				x := val[j] ^ (dep[v] - j)
				if x > res {
					res = x
				}
			}
			ans[i] = res
		}
		for _, w := range g[v] {
			if w != fa {
				dep[w] = dep[v] + 1
				dfs(w, v)
			}
		}
	}
	dfs(1, 0)

	for i := 1; i <= q; i++ {
		Fprintln(out, ans[i])
	}
}

//func main() { cf840E(bufio.NewReader(os.Stdin), os.Stdout) }
