package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF101341A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y, i int }

	var n, d int
	var s []byte
	var ls, rs []pair
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
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
			ls = append(ls, pair{r, l, i})
		} else {
			rs = append(rs, pair{l, r, i})
		}
		d += l - r
	}

	sort.Slice(ls, func(i, j int) bool { return ls[i].x < ls[j].x })
	sort.Slice(rs, func(i, j int) bool { return rs[i].x < rs[j].x })
	f := func(ps []pair) []int {
		ans := []int{}
		s := 0
		for _, p := range ps {
			if s < p.x {
				return nil
			}
			s += p.y - p.x
			ans = append(ans, p.i)
		}
		return ans
	}
	ansL := f(ls)
	ansR := f(rs)
	if d == 0 && ansL != nil && ansR != nil {
		Fprintln(out, "YES")
		for _, v := range ansL {
			Fprint(out, v, " ")
		}
		for i := len(ansR) - 1; i >= 0; i-- {
			Fprint(out, ansR[i], " ")
		}
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF101341A(os.Stdin, os.Stdout) }
