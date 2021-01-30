package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1399E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type nb struct{ to, wt, c int }

	var T, n, v, w, wt, c int
	var s int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		g := make([][]nb, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w, &wt, &c)
			v--
			w--
			g[v] = append(g[v], nb{w, wt, c})
			g[w] = append(g[w], nb{v, wt, c})
		}
		g[0] = append(g[0], nb{-1, 0, 0})
		a, tot := [3][]int64{}, int64(0)
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
						a[e.c] = append(a[e.c], (wt+1)/2*c)
					}
					leaf += c
				}
			}
			return
		}
		f(0, -1)

		for _, b := range a {
			sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
		}
		m := len(a[2])
		s2 := make([]int64, m+1)
		for i, v := range a[2] {
			s2[i+1] = s2[i] + v
		}

		ans := int(1e9)
		a[1] = append(a[1], 0)
		for i, v := range a[1] {
			if j := sort.Search(m+1, func(j int) bool { return tot-s2[j] <= s }); j <= m {
				ans = min(ans, i+2*j)
			}
			tot -= v
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1399E2(os.Stdin, os.Stdout) }
