package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, id int
	Fscan(in, &n)
	g := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
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
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 1 {
				q = append(q, w)
			}
		}
	}

	ids := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		ids[v] = id
		for _, w := range g[v] {
			if w != fa && deg[w] < 2 {
				dfs(w, v)
			}
		}
	}
	for i, d := range deg {
		if d > 1 {
			id++
			dfs(i, -1)
		}
	}

	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &v, &w)
		if ids[v-1] == ids[w-1] {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
