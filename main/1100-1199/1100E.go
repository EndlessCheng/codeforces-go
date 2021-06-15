package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1100E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	es := make([]struct{ v, w, wt int }, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
		es[i].v--
		es[i].w--
	}
	ans := []int{}
	lim := sort.Search(1e9+1, func(lim int) bool {
		g := make([][]int, n)
		d := make([]int, n)
		for _, e := range es {
			if e.wt > lim {
				g[e.v] = append(g[e.v], e.w)
				d[e.w]++
			}
		}
		q, p := []int{}, 0
		for i, d := range d {
			if d == 0 {
				q = append(q, i)
			}
		}
		for ; p < len(q); p++ {
			for _, w := range g[q[p]] {
				if d[w]--; d[w] == 0 {
					q = append(q, w)
				}
			}
		}
		if p == n {
			ans = nil
			order := make([]int, n)
			for i, v := range q {
				order[v] = i
			}
			for i, e := range es {
				if order[e.v] > order[e.w] {
					ans = append(ans, i+1)
				}
			}
		}
		return p == n
	})
	Fprintln(out, lim, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1100E(os.Stdin, os.Stdout) }
