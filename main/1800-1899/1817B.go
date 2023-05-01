package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1817B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		for i, e := range g {
			if len(e) < 4 {
				continue
			}
			vis := make([]bool, n)
			type pair struct{ x, y int }
			path := []pair{}
			onPath := make([]bool, n)
			var dfs func(int, int) bool
			dfs = func(v, fa int) bool {
				vis[v] = true
				for _, w := range g[v] {
					if w == i && fa != i {
						onPath[v] = true
						path = append(path, pair{v, w})
						return true
					}
				}
				for _, w := range g[v] {
					if !vis[w] && dfs(w, v) {
						onPath[v] = true
						path = append(path, pair{v, w})
						return true
					}
				}
				return false
			}
			if !dfs(i, -1) {
				continue
			}
			c := 2
			for _, w := range e {
				if !onPath[w] {
					path = append(path, pair{i, w})
					if c--; c == 0 {
						break
					}
				}
			}
			Fprintln(out, "YES")
			Fprintln(out, len(path))
			for _, e := range path {
				Fprintln(out, e.x+1, e.y+1)
			}
			continue o
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1817B(os.Stdin, os.Stdout) }
