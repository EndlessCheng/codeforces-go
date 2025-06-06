package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2110D(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		type nb struct{ to, wt int }
		g := make([][]nb, n)
		for range m {
			var v, w, wt int
			Fscan(in, &v, &w, &wt)
			g[v-1] = append(g[v-1], nb{w - 1, wt})
		}

		f := make([]int, n)
		ans := sort.Search(1e9+1, func(mx int) bool {
			clear(f)
			for v, fv := range f {
				if v == 0 || fv > 0 {
					fv = min(fv+b[v], mx)
				}
				for _, e := range g[v] {
					w := e.to
					if e.wt <= fv {
						f[w] = max(f[w], fv)
					}
				}
			}
			return f[n-1] > 0
		})
		if ans > 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2110D(bufio.NewReader(os.Stdin), os.Stdout) }
