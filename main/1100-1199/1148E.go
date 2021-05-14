package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1148E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ v, i int }

	var n int
	Fscan(in, &n)
	s := make([]pair, n)
	for i := range s {
		Fscan(in, &s[i].v)
		s[i].i = i + 1
	}
	sort.Slice(s, func(i, j int) bool { return s[i].v < s[j].v })
	t := make([]int, n)
	for i := range t {
		Fscan(in, &t[i])
	}
	sort.Ints(t)

	ans := [][3]int{}
	q := []pair{}
	for i, x := range s {
		d := x.v - t[i]
		if d < 0 {
			q = append(q, pair{-d, x.i})
			continue
		}
		for d > 0 {
			if len(q) == 0 {
				Fprint(out, "NO")
				return
			}
			y := &q[len(q)-1]
			if d < y.v {
				y.v -= d
				ans = append(ans, [3]int{y.i, x.i, d})
				break
			}
			q = q[:len(q)-1]
			d -= y.v
			ans = append(ans, [3]int{y.i, x.i, y.v})
		}
	}
	if len(q) > 0 {
		Fprint(out, "NO")
	} else {
		Fprintln(out, "YES")
		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0], p[1], p[2])
		}
	}
}

//func main() { CF1148E(os.Stdin, os.Stdout) }
