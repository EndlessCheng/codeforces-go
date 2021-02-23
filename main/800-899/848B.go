package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF848B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ p, i int }

	var n, w, h, g, p, t int
	Fscan(in, &n, &w, &h)
	var vs, hs [2e5][]pair
	for i := 0; i < n; i++ {
		if Fscan(in, &g, &p, &t); g == 1 {
			vs[p+1e5-t] = append(vs[p+1e5-t], pair{p, i})
		} else {
			hs[p+1e5-t] = append(hs[p+1e5-t], pair{p, i})
		}
	}

	ans := make([][2]int, n)
	for i, vs := range vs[:] {
		hs := hs[i]
		sort.Slice(vs, func(i, j int) bool { return vs[i].p < vs[j].p })
		sort.Slice(hs, func(i, j int) bool { return hs[i].p > hs[j].p })
		for _, p := range append(hs, vs...) {
			if len(vs) > 0 {
				ans[p.i] = [2]int{vs[0].p, h}
				vs = vs[1:]
			} else {
				ans[p.i] = [2]int{w, hs[0].p}
				hs = hs[1:]
			}
		}
	}
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF848B(os.Stdin, os.Stdout) }
