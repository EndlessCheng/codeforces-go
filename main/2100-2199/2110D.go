package main

import (
	. "fmt"
	"io"
	"slices"
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
		deg0 := make([]int, n)
		for range m {
			var v, w, wt int
			Fscan(in, &v, &w, &wt)
			g[v-1] = append(g[v-1], nb{w - 1, wt})
			deg0[w-1]++
		}

		q := []int{}
		for i := 1; i < n; i++ {
			if deg0[i] == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				w := e.to
				if deg0[w]--; deg0[w] == 0 {
					q = append(q, w)
				}
			}
		}

		ans := sort.Search(1e9+1, func(mx int) bool {
			deg := slices.Clone(deg0)
			f := make([]int, n)
			q := []int{0}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				if v == n-1 {
					return f[v] > 0
				}
				if v == 0 || f[v] > 0 {
					f[v] = min(f[v]+b[v], mx)
				}
				for _, e := range g[v] {
					w := e.to
					if e.wt <= f[v] {
						f[w] = max(f[w], f[v])
					}
					if deg[w]--; deg[w] == 0 {
						q = append(q, w)
					}
				}
			}
			return false
		})
		if ans > 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2110D(bufio.NewReader(os.Stdin), os.Stdout) }
