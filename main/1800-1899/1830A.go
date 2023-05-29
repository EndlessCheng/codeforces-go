package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1830A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type nb struct{ to, i int }
		g := make([][]nb, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], nb{w, i})
			g[w] = append(g[w], nb{v, i})
		}
		
		var dfs func(int, int) int
		dfs = func(v, i int) (mx int) {
			for _, e := range g[v] {
				if e.i != i {
					res := dfs(e.to, e.i)
					if e.i < i {
						res++
					}
					mx = max(mx, res)
				}
			}
			return
		}
		Fprintln(out, dfs(0, -1)+1)
	}
}

//func main() { CF1830A(os.Stdin, os.Stdout) }
