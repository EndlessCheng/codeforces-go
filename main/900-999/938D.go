package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf938D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt * 2})
		g[w] = append(g[w], nb{v, wt * 2})
	}
	dis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &wt)
		g[0] = append(g[0], nb{i, wt})
		dis[i] = 1e18
	}

	h := hp38{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair38)
		v := p.v
		if p.dis > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			newD := p.dis + e.wt
			if newD < dis[w] {
				dis[w] = newD
				heap.Push(&h, pair38{newD, w})
			}
		}
	}
	for _, v := range dis[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { cf938D(os.Stdin, os.Stdout) }
type pair38 struct{ dis, v int }
type hp38 []pair38
func (h hp38) Len() int           { return len(h) }
func (h hp38) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp38) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp38) Push(v any)        { *h = append(*h, v.(pair38)) }
func (h *hp38) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
