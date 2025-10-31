package main

import (
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1608(in io.Reader, out io.Writer) {
	const inf int = 1e18
	var n, m int
	Fscan(in, &n, &m)
	minE := map[[2]int]int{}
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		p := [2]int{v - 1, w - 1}
		if res, ok := minE[p]; !ok || wt < res {
			minE[p] = wt
		}
	}
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for p, wt := range minE {
		g[p[0]] = append(g[p[0]], nb{p[1], wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[0] = 0
	f := make([]int, n)
	f[0] = 1
	h := hp1608{{}}
	for len(h) > 0 {
		p := h.pop()
		v := p.v
		if p.dis > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			newD := p.dis + e.wt
			if newD < dis[w] {
				dis[w] = newD
				f[w] = f[v]
				h.push(vdPair1608{newD, w})
			} else if newD == dis[w] {
				f[w] += f[v]
			}
		}
	}
	if dis[n-1] == inf {
		Fprintln(out, "No answer")
	} else {
		Fprintln(out, dis[n-1], f[n-1])
	}
}

//func main() { p1608(bufio.NewReader(os.Stdin), os.Stdout) }

type vdPair1608 struct{ dis, v int }
type hp1608 []vdPair1608

func (h hp1608) Len() int           { return len(h) }
func (h hp1608) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp1608) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp1608) Push(v any)        { *h = append(*h, v.(vdPair1608)) }
func (h *hp1608) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp1608) push(v vdPair1608) { heap.Push(h, v) }
func (h *hp1608) pop() vdPair1608   { return heap.Pop(h).(vdPair1608) }
