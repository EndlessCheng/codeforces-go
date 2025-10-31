package main

import (
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2865(in io.Reader, out io.Writer) {
	const inf int = 1e18
	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[0] = 0
	dis2 := make([]int, n)
	for i := range dis2 {
		dis2[i] = inf
	}
	h := hp2865{{}}
	for len(h) > 0 {
		p := h.pop()
		v, d := p.v, p.d
		if dis2[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			newD := d + e.wt
			if newD < dis[w] {
				dis2[w] = dis[w]
				dis[w] = newD
				h.push(pair2865{w, newD})
			} else if dis[w] < newD && newD < dis2[w] {
				dis2[w] = newD
				h.push(pair2865{w, newD})
			}
		}
	}
	Fprint(out, dis2[n-1])
}

//func main() { p2865(bufio.NewReader(os.Stdin), os.Stdout) }

type pair2865 struct{ v, d int }
type hp2865 []pair2865

func (h hp2865) Len() int              { return len(h) }
func (h hp2865) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp2865) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp2865) Push(v interface{})   { *h = append(*h, v.(pair2865)) }
func (h *hp2865) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp2865) push(v pair2865)      { heap.Push(h, v) }
func (h *hp2865) pop() pair2865        { return heap.Pop(h).(pair2865) }
