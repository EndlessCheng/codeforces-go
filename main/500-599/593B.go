package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF593B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x1, x2, k, b int64
	Fscan(in, &n, &x1, &x2)
	type pair struct{ l, r int64 }
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &k, &b)
		ps[i].l = k*x1 + b
		ps[i].r = k*x2 + b
	}
	sort.Slice(ps, func(i, j int) bool {
		a, b := ps[i], ps[j]
		return a.l < b.l || a.l == b.l && a.r < b.r
	})
	for i, p := range ps[:n-1] {
		if p.r > ps[i+1].r {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() {
//	CF593B(os.Stdin, os.Stdout)
//}
