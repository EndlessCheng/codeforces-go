package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1225F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	pa := make([]int, n)
	g := make([][]int, n)
	dep := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &pa[i])
		g[pa[i]] = append(g[pa[i]], i)
		dep[i] = dep[pa[i]] + 1
	}
	dn := make([]int, n)
	idx := make([]int, n)
	for i := n - 1; i >= 1; i-- {
		if dn[pa[i]] < dn[i]+1 {
			dn[pa[i]] = dn[i] + 1
			idx[pa[i]] = i
		}
	}

	dfn := make([]int, n)
	tot := 0
	var dfs func(int)
	dfs = func(v int) {
		dfn[tot] = v
		tot++
		for _, w := range g[v] {
			if w != idx[v] {
				dfs(w)
			}
		}
		if idx[v] != 0 {
			dfs(idx[v])
		}
	}
	dfs(0)

	ans := 0
	for i := 1; i < n; i++ {
		ans += dep[dfn[i-1]] - dep[pa[dfn[i]]]
	}

	for _, v := range dfn {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	Fprintln(out, ans)
	for i := 1; i < n; i++ {
		for range dep[dfn[i-1]] - dep[pa[dfn[i]]] {
			Fprint(out, dfn[i], " ")
		}
	}
}

//func main() { cf1225F(bufio.NewReader(os.Stdin), os.Stdout) }
