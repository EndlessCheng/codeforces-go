package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1689C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		g[1] = []int{0}
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			if len(g[v]) <= 2 {
				return len(g[v])
			}
			mx := int(1e9)
			for _, w := range g[v] {
				if w != fa {
					mx = min(mx, dfs(w, v)+2)
				}
			}
			return mx
		}
		Fprintln(out, n-dfs(1, 0))
	}
}

//func main() { cf1689C(bufio.NewReader(os.Stdin), os.Stdout) }
