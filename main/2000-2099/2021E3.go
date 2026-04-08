package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2021E3(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &p)
		cnt := make([]int, n)
		for range p {
			var v int
			Fscan(in, &v)
			v--
			cnt[v] = 1
		}
		type edge struct{ wt, u, v int }
		es := make([]edge, m)
		for i := range es {
			var u, v, w int
			Fscan(in, &u, &v, &w)
			es[i] = edge{w, u - 1, v - 1}
		}
		slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

		fa := make([]int, n)
		for i := range n {
			fa[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}

		ws := make([]int, n)
		d := make([]int, 0, p-1)
		for _, e := range es {
			u := find(e.u)
			v := find(e.v)
			if u == v {
				continue
			}
			if cnt[u] < cnt[v] {
				u, v = v, u
			}
			fa[v] = u
			if cnt[v] == 0 {
				continue
			}
			cu := cnt[u]
			cv := cnt[v]
			cnt[u] += cnt[v]
			nw := min(ws[u]+e.wt*cv, ws[v]+e.wt*cu)
			d = append(d, ws[u]+ws[v]-nw)
			ws[u] = nw
		}

		s := ws[find(0)]
		slices.Sort(d)
		for i := range n {
			Fprint(out, s, " ")
			if i < len(d) {
				s += d[i]
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2021E3(bufio.NewReader(os.Stdin), os.Stdout) }
