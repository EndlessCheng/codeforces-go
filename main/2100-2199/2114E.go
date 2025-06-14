package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2114E(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := make([]any, n)
		var dfs func(int, int, int, int)
		dfs = func(v, fa, pos, neg int) {
			pos, neg = neg+a[v], max(pos-a[v], 0)
			ans[v] = pos
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, pos, neg)
				}
			}
		}
		dfs(0, -1, 0, 0)
		Fprintln(out, ans...)
	}
}

//func main() { cf2114E(bufio.NewReader(os.Stdin), os.Stdout) }
