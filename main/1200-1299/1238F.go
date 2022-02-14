package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1238F(_r io.Reader, _w io.Writer) {
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
		g := make([][]int, n+1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := 0
		var f func(int, int) int
		f = func(v, fa int) int {
			mx := 0
			for _, w := range g[v] {
				if w != fa {
					d := f(w, v)
					ans = max(ans, mx+d+len(g[v])-1)
					mx = max(mx, d)
				}
			}
			return mx + len(g[v]) - 1
		}
		f(1, 0)
		Fprintln(out, ans+2)
	}
}

//func main() { CF1238F(os.Stdin, os.Stdout) }
