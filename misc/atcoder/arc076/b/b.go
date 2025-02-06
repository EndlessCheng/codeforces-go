package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	type pair struct{ v, i int }
	a := make([]pair, n)
	b := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v, &b[i].v)
		a[i].i = i
		b[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	sort.Slice(b, func(i, j int) bool { return b[i].v < b[j].v })

	type edge struct{ v, w, wt int }
	es := make([]edge, 0, n*2-2)
	for i := 1; i < n; i++ {
		es = append(es, edge{a[i-1].i, a[i].i, a[i].v - a[i-1].v})
		es = append(es, edge{b[i-1].i, b[i].i, b[i].v - b[i-1].v})
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	for _, e := range es {
		v := find(e.v)
		w := find(e.w)
		if v != w {
			fa[v] = w
			ans += e.wt
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
