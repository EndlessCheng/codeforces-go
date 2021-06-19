package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		e := make([]struct{ t, v, w int }, m)
		g := make([][]int, n)
		d := make([]int, n)
		for i := range e {
			Fscan(in, &e[i].t, &e[i].v, &e[i].w)
			e[i].v--
			e[i].w--
			if e[i].t > 0 {
				g[e[i].v] = append(g[e[i].v], e[i].w)
				d[e[i].w]++
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
			continue
		}
		Fprintln(out, "YES")
		order := make([]int, n)
		for i, v := range q {
			order[v] = i
		}
		for _, e := range e {
			v, w := e.v, e.w
			if e.t == 0 && order[v] > order[w] {
				v, w = w, v
			}
			Fprintln(out, v+1, w+1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
