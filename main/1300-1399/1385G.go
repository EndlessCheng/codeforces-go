package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1385G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type nb struct{ w, i int }

	var T, n, w int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]nb, n+1)
		vis := make([]bool, n+1)
		for i, v := range a {
			Fscan(in, &w)
			g[v] = append(g[v], nb{w, i + 1})
			g[w] = append(g[w], nb{v, -(i + 1)})
			if v == w {
				vis[v] = true
			}
		}
		for _, es := range g[1:] {
			if len(es) != 2 {
				Fprintln(out, -1)
				continue o
			}
		}
		ans := []int{}
		for i := 1; i <= n; i++ {
			if vis[i] {
				continue
			}
			if g[i][0].w == g[i][1].w {
				if g[i][0].i > 0 == (g[i][1].i > 0) {
					if g[i][0].i < 0 {
						g[i][0].i = -g[i][0].i
					}
					ans = append(ans, g[i][0].i)
				}
				vis[g[i][0].w] = true
				continue
			}
			var i1, i2 []int
			pre, v := -1, i
			for {
				vis[v] = true
				for _, e := range g[v] {
					if e.w != pre {
						pre, v = v, e.w
						if e.i > 0 {
							i1 = append(i1, e.i)
						} else {
							i2 = append(i2, -e.i)
						}
						break
					}
				}
				if v == i {
					break
				}
			}
			if len(i1) < len(i2) {
				ans = append(ans, i1...)
			} else {
				ans = append(ans, i2...)
			}
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1385G(os.Stdin, os.Stdout) }
