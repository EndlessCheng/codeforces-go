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

	var n, q, v, w, p, x int
	Fscan(in, &n, &q)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	a := make([]int, n+1)
	for ; q > 0; q-- {
		Fscan(in, &p, &x)
		a[p] += x
	}

	var dfs func(int, int)
	dfs = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa {
				a[w] += a[v]
				dfs(w, v)
			}
		}
	}
	dfs(1, 0)
	for _, v := range a[1:] {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
