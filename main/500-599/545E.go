package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf545E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, st int
	Fscan(in, &n, &m)
	type edge struct{ to, wt, i int }
	g := make([][]edge, n)
	for i := 1; i <= m; i++ {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], edge{w, wt, i})
		g[w] = append(g[w], edge{v, wt, i})
	}
	Fscan(in, &st)
	st--

	type pair struct{ wt, i int }
	from := make([]pair, n)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1e18
	}
	dis[st] = 0
	h := hp45{{0, st}}
	for len(h) > 0 {
		p := heap.Pop(&h).(vd45)
		v := p.v
		d := p.dis
		if d > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w, wt := e.to, e.wt
			newD := d + wt
			if newD < dis[w] {
				dis[w] = newD
				heap.Push(&h, vd45{newD, w})
				from[w] = pair{wt, e.i}
			} else if newD == dis[w] {
				from[w] = pair{wt, e.i}
			}
		}
	}

	ans := 0
	for _, p := range from {
		ans += p.wt
	}
	Fprintln(out, ans)
	for i, p := range from {
		if i != st {
			Fprint(out, p.i, " ")
		}
	}
}

//func main() { cf545E(bufio.NewReader(os.Stdin), os.Stdout) }

type vd45 struct{ dis, v int }
type hp45 []vd45
func (h hp45) Len() int           { return len(h) }
func (h hp45) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp45) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp45) Push(v any)        { *h = append(*h, v.(vd45)) }
func (h *hp45) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
