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

	var n, v, w, q int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	Fscan(in, &q)
	type pair struct{ i, k int }
	qs := make([][]pair, n+1)
	for i := 0; i < q; i++ {
		Fscan(in, &v, &w)
		qs[v] = append(qs[v], pair{i, w})
	}

	ans := make([]int, q)
	for i := range ans {
		ans[i] = -1
	}
	nodes := make([]int, n)
	for i, rt := 0, 1; i < 3; i++ {
		mx := -1
		var dfs func(int, int, int)
		dfs = func(v, fa, d int) {
			if d > mx {
				mx = d
				rt = v
			}
			nodes[d] = v
			for _, p := range qs[v] {
				if d >= p.k {
					ans[p.i] = nodes[d-p.k]
				}
			}
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, d+1)
				}
			}
		}
		dfs(rt, 0, 0)
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
