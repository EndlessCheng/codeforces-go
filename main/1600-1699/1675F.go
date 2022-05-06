package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1675F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, x, y, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x, &y)
		need := make([]bool, n+1)
		for ; k > 0; k-- {
			Fscan(in, &v)
			need[v] = true
		}
		need[x] = true
		need[y] = true
		g := make([][]int, n+1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := n*2 - 2
		var f func(v, fa, d int) bool
		f = func(v, fa, d int) bool {
			if v == y {
				ans -= d
			}
			found := need[v]
			for _, w := range g[v] {
				if w != fa && f(w, v, d+1) {
					found = true
				}
			}
			if !found {
				ans -= 2
			}
			return found
		}
		f(x, 0, 0)
		Fprintln(out, ans)
	}
}

//func main() { CF1675F(os.Stdin, os.Stdout) }
