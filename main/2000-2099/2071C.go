package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2071C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, t, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &t, &t)
		g := make([][]int, n+1)
		for range n - 1 {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		dep := make([][]int, n)
		var dfs func(int, int, int)
		dfs = func(v, fa, d int) {
			dep[d] = append(dep[d], v)
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, d+1)
				}
			}
		}
		dfs(t, 0, 0)

		for i := n - 1; i >= 0; i-- {
			for _, v := range dep[i] {
				Fprint(out, v, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2071C(bufio.NewReader(os.Stdin), os.Stdout) }
