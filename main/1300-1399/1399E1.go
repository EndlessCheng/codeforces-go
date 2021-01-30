package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1399E1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type nb struct{ to, wt int }

	var T, n, v, w, wt int
	var s int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		g := make([][]nb, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w, &wt)
			v--
			w--
			g[v] = append(g[v], nb{w, wt})
			g[w] = append(g[w], nb{v, wt})
		}
		g[0] = append(g[0], nb{-1, 0})
		a, tot := []int64{}, int64(0)
		var f func(v, fa int) int64
		f = func(v, fa int) (leaf int64) {
			if len(g[v]) == 1 {
				return 1
			}
			for _, e := range g[v] {
				if e.to != fa {
					c := f(e.to, v)
					wt := int64(e.wt)
					tot += wt * c
					for ; wt > 0; wt >>= 1 {
						a = append(a, (wt+1)/2*c)
					}
					leaf += c
				}
			}
			return
		}
		f(0, -1)
		sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
		for i, v := range a {
			if tot <= s {
				Fprintln(out, i)
				break
			}
			tot -= v
		}
	}
}

//func main() { CF1399E1(os.Stdin, os.Stdout) }
