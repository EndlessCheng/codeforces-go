package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1385E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		es := make([]struct{ v, w, tp int }, m)
		g := make([][]int, n)
		d := make([]int, n)
		for i := range es {
			Fscan(in, &es[i].tp, &es[i].v, &es[i].w)
			es[i].v--
			es[i].w--
			if es[i].tp == 1 {
				g[es[i].v] = append(g[es[i].v], es[i].w)
				d[es[i].w]++
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
		if p < n {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			t := make([]int, n)
			for i, v := range q {
				t[v] = i
			}
			for _, e := range es {
				v, w := e.v, e.w
				if e.tp == 0 && t[v] > t[w] {
					v, w = w, v
				}
				Fprintln(out, v+1, w+1)
			}
		}
	}
}

//func main() { CF1385E(os.Stdin, os.Stdout) }
