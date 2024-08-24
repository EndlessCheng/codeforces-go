package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

func p1144(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1e18
	}
	dis[0] = 0
	dp := make([]int, n)
	dp[0] = 1
	h := hp1144{{}}
	for len(h) > 0 {
		p := h.pop()
		v := p.v
		if p.d > dis[v] {
			continue
		}
		for _, w := range g[v] {
			if newD := dis[v] + 1; newD < dis[w] {
				dis[w] = newD
				dp[w] = dp[v]
				h.push(vd1144{w, newD})
			} else if newD == dis[w] {
				dp[w] = (dp[w] + dp[v]) % (1e5 + 3)
			}
		}
	}
	for _, v := range dp {
		Fprintln(out, v)
	}
}

//func main() { p1144(os.Stdin, os.Stdout) }

type vd1144 struct{ v, d int }
type hp1144 []vd1144

func (h hp1144) Len() int              { return len(h) }
func (h hp1144) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp1144) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp1144) Push(v interface{})   { *h = append(*h, v.(vd1144)) }
func (h *hp1144) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp1144) push(v vd1144)        { heap.Push(h, v) }
func (h *hp1144) pop() vd1144          { return heap.Pop(h).(vd1144) }
