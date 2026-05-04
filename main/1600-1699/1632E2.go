package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1632E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		dep := make([]int, n+1)
		dep[0] = -1
		mx := make([]int, n+1)
		f := make([]int, n+2)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			mx[v] = dep[fa] + 1
			dep[v] = dep[fa] + 1
			m := mx[v]
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				dfs(w, v)
				if mx[v] < mx[w] {
					m = mx[v]
					mx[v] = mx[w]
				} else {
					m = max(m, mx[w])
				}
			}
			if m > 0 {
				f[m-1] = max(f[m-1], mx[v]+m-dep[v]*2)
			}
		}
		dfs(1, 0)
		d1 := mx[1]

		for i := n; i > 0; i-- {
			f[i] = max(f[i], f[i+1])
		}

		j := 0
		for i := 1; i <= n; i++ {
			for j < d1 && (f[j]+1)/2+i > j {
				j++
			}
			Fprint(out, j, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1632E2(bufio.NewReader(os.Stdin), os.Stdout) }
