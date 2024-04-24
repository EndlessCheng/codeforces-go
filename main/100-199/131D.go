package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf131D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	g := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}
	rg := make([][]int, n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			deg[w]--
			if deg[w] > 0 {
				rg[w] = append(rg[w], v)
				if deg[w] == 1 {
					q = append(q, w)
				}
			}
		}
	}

	ans := make([]int, n)
	var dfs func(int)
	dfs = func(v int) {
		for _, w := range rg[v] {
			ans[w] = ans[v] + 1
			dfs(w)
		}
	}
	for i, v := range ans {
		if v == 0 {
			dfs(i)
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { cf131D(os.Stdin, os.Stdout) }
