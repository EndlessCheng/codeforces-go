package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1278D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, l, r, c int
	Fscan(in, &n)
	type event struct{ p, d, i int }
	es := make([]event, 0, 2*n)
	for i := 0; i < n; i++ {
		Fscan(in, &l, &r)
		es = append(es, event{l, 1, i}, event{r, 0, i})
	}
	sort.Slice(es, func(i, j int) bool { return es[i].p < es[j].p })

	g := make([][]int, n)
	s := []event{}
o:
	for _, e := range es {
		if e.d > 0 {
			s = append(s, e)
		} else {
			for i := len(s) - 1; ; i-- {
				e2 := s[i]
				if e2.i == e.i {
					s = append(s[:i], s[i+1:]...)
					break
				}
				if c++; c >= n {
					break o
				}
				g[e.i] = append(g[e.i], e2.i)
				g[e2.i] = append(g[e2.i], e.i)
			}
		}
	}
	if c == n-1 {
		c++
		vis := make([]bool, n)
		var f func(int)
		f = func(v int) {
			vis[v] = true
			c--
			for _, w := range g[v] {
				if !vis[w] {
					f(w)
				}
			}
			return
		}
		f(0)
		if c == 0 {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF1278D(os.Stdin, os.Stdout) }
