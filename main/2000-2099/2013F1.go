package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2013F1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		Fscan(in, &t, &t)
		t--

		vis := make([]bool, n)
		var dfs func(int, int) bool
		dfs = func(v, nxt int) bool {
			for _, w := range g[v] {
				if !vis[w] {
					vis[w] = true
					if !dfs(nxt, w) {
						vis[w] = false
						return true
					}
					vis[w] = false
				}
			}
			return false
		}

		vis[0] = true
		vis[t] = true
		if dfs(0, t) {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { cf2013F1(bufio.NewReader(os.Stdin), os.Stdout) }
