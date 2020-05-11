package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	var n, d int
	var s []byte
	var ls, rs []pair
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		l, r := 0, 0
		for _, b := range s {
			if b == '(' {
				l++
			} else if l > 0 {
				l--
			} else {
				r++
			}
		}
		if r < l {
			ls = append(ls, pair{r, l})
		} else {
			rs = append(rs, pair{l, r})
		}
		d += l - r
	}

	sort.Slice(ls, func(i, j int) bool { return ls[i].x < ls[j].x })
	sort.Slice(rs, func(i, j int) bool { return rs[i].x < rs[j].x })
	f := func(ps []pair) bool {
		s := 0
		for _, p := range ps {
			if s < p.x {
				return false
			}
			s += p.y - p.x
		}
		return true
	}
	if d == 0 && f(ls) && f(rs) {
		Fprint(_w, "Yes")
	} else {
		Fprint(_w, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }
