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
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, wt int
	var s int64
	Fscan(in, &n, &m, &s)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}

	type edge struct{ v, w, wt, i int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		es[i] = edge{v - 1, w - 1, wt, i + 1}
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt > es[j].wt })
	for i, e := range es {
		if fv, fw := f(e.v), f(e.w); fv != fw {
			fa[fv] = fw
			es[i].i = 0
		}
	}
	id := []interface{}{}
	for i := len(es) - 1; i >= 0; i-- {
		if e := es[i]; e.i > 0 && int64(e.wt) <= s {
			s -= int64(e.wt)
			id = append(id, e.i)
		}
	}
	Fprintln(out, len(id))
	Fprint(out, id...)
}

func main() { run(os.Stdin, os.Stdout) }
