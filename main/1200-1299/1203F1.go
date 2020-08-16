package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1203F1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	var n, r, a, b int
	var ls, rs []pair
	Fscan(in, &n, &r)
	ls = append(ls, pair{0, r})
	d := r
	for ; n > 0; n-- {
		Fscan(in, &a, &b)
		if b >= 0 {
			ls = append(ls, pair{a, b})
		} else {
			rs = append(rs, pair{a + b, a})
		}
		d += b
	}
	if d < 0 {
		Fprint(_w, "NO")
		return
	}
	rs = append(rs, pair{0, d})

	sort.Slice(ls, func(i, j int) bool { return ls[i].x < ls[j].x })
	sort.Slice(rs, func(i, j int) bool { return rs[i].x < rs[j].x })
	f := func(ps []pair) bool {
		s := 0
		for _, p := range ps {
			if s < p.x {
				return false
			}
			s += p.y
		}
		return true
	}
	if f(ls) && f(rs) { // f(rs) 相当于从结局倒着思考
		Fprint(_w, "YES")
	} else {
		Fprint(_w, "NO")
	}
}

//func main() { CF1203F1(os.Stdin, os.Stdout) }
