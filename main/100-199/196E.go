package main

import (
	"container/heap"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf196E(in io.Reader, out io.Writer) {
	var n, m, k, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range m {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
		es[i] = edge{v, w, wt}
	}

	dis := make([]int, n)
	from := make([]int, n)
	for i := range dis {
		dis[i] = 1e18
		from[i] = i
	}
	Fscan(in, &k)
	h := make(hp96, k)
	for i := range k {
		Fscan(in, &v)
		v--
		dis[v] = 0
		h[i].v = v
	}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair96)
		v := p.v
		d := p.d
		if d > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			newD := d + e.wt
			if newD < dis[w] {
				dis[w] = newD
				from[w] = from[v] // 易错点
				heap.Push(&h, pair96{newD, w})
			}
		}
	}

	ans := dis[0]
	for i, e := range es {
		es[i].wt += dis[e.v] + dis[e.w]
	}
	slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

	fa := dis
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
		v := find(from[e.v])
		w := find(from[e.w])
		if v != w {
			ans += e.wt
			fa[v] = w
		}
	}
	Fprint(out, ans)
}

//func main() { cf196E(bufio.NewReader(os.Stdin), os.Stdout) }
type pair96 struct{ d, v int }
type hp96 []pair96
func (h hp96) Len() int           { return len(h) }
func (h hp96) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp96) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp96) Push(v any)        { *h = append(*h, v.(pair96)) }
func (h *hp96) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
