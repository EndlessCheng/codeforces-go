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

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n+1)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				Fprintln(out, v, w)
				dfs(w)
			}
		}
	}
	dfs(1)

	vis = make([]bool, n+1)
	vis[1] = true
	q := []int{1}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				q = append(q, w)
				Fprintln(out, v, w)
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
