package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

func cf1851G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, q, v, w, e int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type pair struct{ v, i int }
		h := make([]pair, n)
		for i := range h {
			Fscan(in, &h[i].v)
			h[i].i = i
		}

		g := make([][]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			v--
			w--
			if h[v].v < h[w].v {
				v, w = w, v
			}
			g[v] = append(g[v], w)
		}

		Fscan(in, &q)
		type data struct{ e, v, w, qid int }
		qs := make([]data, q)
		for i := range qs {
			Fscan(in, &v, &w, &e)
			v--
			w--
			qs[i] = data{e + h[v].v, v, w, i}
		}
		slices.SortFunc(qs, func(a, b data) int { return a.e - b.e })
		slices.SortFunc(h, func(a, b pair) int { return a.v - b.v })

		fa := make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}

		ans := make([]bool, q)
		i := 0
		for _, q := range qs {
			for ; i < n && h[i].v <= q.e; i++ {
				f := find(h[i].i)
				for _, w := range g[h[i].i] {
					fa[find(w)] = f
				}
			}
			ans[q.qid] = find(q.v) == find(q.w)
		}
		for _, b := range ans {
			if b {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1851G(bufio.NewReader(os.Stdin), os.Stdout) }
