package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1223E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type nb struct{ to, wt int }
	type pair struct{ c, nc int64 } // 选 v-w，不选 v-w

	var T, n, k, v, w, wt int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := make([][]nb, n+1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w, &wt)
			g[v] = append(g[v], nb{w, wt})
			g[w] = append(g[w], nb{v, wt})
		}
		var f func(int, int) (int64, int64)
		f = func(v, fa int) (int64, int64) {
			sum := int64(0)
			a := []pair{}
			for _, e := range g[v] {
				if w := e.to; w != fa {
					nc, c := f(w, v)
					c += int64(e.wt)
					if nc >= c {
						sum += nc
					} else {
						a = append(a, pair{c, nc})
					}
				}
			}
			if len(a) < k {
				for _, p := range a {
					sum += p.c
				}
				return sum, sum
			}
			sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.c-a.nc > b.c-b.nc })
			for _, p := range a[:k-1] {
				sum += p.c
			}
			for _, p := range a[k:] {
				sum += p.nc
			}
			p := a[k-1]
			return sum + p.c, sum + p.nc
		}
		ans, _ := f(1, 0)
		Fprintln(out, ans)
	}
}

//func main() { CF1223E(os.Stdin, os.Stdout) }
