package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1442E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &c[i])
		}
		g := make([][]int, n+1)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		var end, maxD int
		var dfs func(int, int, int, int)
		dfs = func(v, fa, k, dep int) {
			if dep > maxD {
				end = v
				maxD = dep
			}
			for _, w := range g[v] {
				if w != fa {
					nk := k
					if c[w] != 0 {
						nk = c[w]
					}
					wt := 0
					if c[w] != 0 && k != c[w] {
						wt = 1
					}
					dfs(w, v, nk, dep+wt)
				}
			}
		}

		maxD = -1
		if c[1] != 0 {
			dfs(1, 0, c[1], 1)
		} else {
			dfs(1, 0, c[1], 0)
		}

		maxD = -1
		if c[end] != 0 {
			dfs(end, 0, c[end], 1)
		} else {
			dfs(end, 0, c[end], 0)
		}

		Fprintln(out, maxD/2+1)
	}
}

//func main() { cf1442E(bufio.NewReader(os.Stdin), os.Stdout) }
