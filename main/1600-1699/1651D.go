package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1651D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, x, y int
	Fscan(in, &n)
	s := make(map[pair]int, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &x, &y)
		s[pair{x, y}] = i
	}

	ans := make([]pair, n)
	type p2 struct {
		i, j int
		pair
	}
	q := []p2{}
	for v, i := range s {
		for _, d := range dir4 {
			w := pair{v.x + d.x, v.y + d.y}
			if s[w] == 0 {
				ans[i-1] = w
				q = append(q, p2{v.x, v.y, w})
				break
			}
		}
	}
	for _, p := range q {
		delete(s, pair{p.i, p.j})
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, d := range dir4 {
			w := pair{v.i + d.x, v.j + d.y}
			if i := s[w]; i > 0 {
				delete(s, w)
				ans[i-1] = v.pair
				q = append(q, p2{w.x, w.y, v.pair})
			}
		}
	}
	for _, p := range ans {
		Fprintln(out, p.x, p.y)
	}
}

//func main() { CF1651D(os.Stdin, os.Stdout) }
