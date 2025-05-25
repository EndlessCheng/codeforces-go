package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2062D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ l, r int }, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
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

		inc := 0
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			p := a[v]
			mx := 0
			for _, w := range g[v] {
				if w != fa {
					fw := dfs(w, v)
					inc += max(fw-p.r, 0)
					mx = max(mx, fw)
				}
			}
			return max(p.l, min(mx, p.r))
		}
		Fprintln(out, dfs(0, -1)+inc)
	}
}

//func main() { cf2062D(bufio.NewReader(os.Stdin), os.Stdout) }
