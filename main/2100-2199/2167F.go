package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2167F(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := make([][]int, n+1)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := n
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			sz := 1
			for _, w := range g[v] {
				if w != fa {
					sz += dfs(w, v)
				}
			}
			if sz >= k {
				ans += n - sz
			}
			if n-sz >= k {
				ans += sz
			}
			return sz
		}
		dfs(1, 0)
		Fprintln(out, ans)
	}
}

//func main() { cf2167F(bufio.NewReader(os.Stdin), os.Stdout) }
