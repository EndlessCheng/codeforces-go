package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1976F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		res := make([]int, n+1)
		f := make([]int, n+1)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			q := 0
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v)
					if f[w] >= f[q] {
						q = w
					}
				}
			}
			f[v] = f[q] + 1
			for _, w := range g[v] {
				if w != fa && w != q {
					res[w] = f[w]
				}
			}
		}
		dfs(1, 0)

		res[1] = f[1] - 1
		slices.SortFunc(res[1:], func(a, b int) int { return b - a })
		for i := 1; i <= n; i++ {
			res[i] += res[i-1]
		}

		for i := 1; i < n; i++ {
			Fprint(out, n-1-res[min(i*2-1, n)], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1976F(bufio.NewReader(os.Stdin), os.Stdout) }
