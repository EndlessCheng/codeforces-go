package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair15 struct {
	v int
	d int64
}
type hp15 []pair15

func (h hp15) Len() int              { return len(h) }
func (h hp15) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp15) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp15) Push(v interface{})   { *h = append(*h, v.(pair15)) }
func (h *hp15) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp15) push(v pair15)        { heap.Push(h, v) }
func (h *hp15) pop() pair15          { return heap.Pop(h).(pair15) }

func CF715B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, s, t, v, w int
	var l, wt, diff int64
	Fscan(in, &n, &m, &l, &s, &t)
	type nb struct {
		to, rid int
		wt      int64
		del     bool
	}
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		del := wt == 0
		if del {
			wt = 1
		}
		g[v] = append(g[v], nb{w, len(g[w]), wt, del})
		g[w] = append(g[w], nb{v, len(g[v]) - 1, wt, del})
	}

	dis := make([][2]int64, n)
	for i := range dis {
		dis[i] = [2]int64{1e18, 1e18}
	}
	dis[s] = [2]int64{}
	dij := func(k int) {
		h := hp15{{s, 0}}
		for len(h) > 0 {
			p := h.pop()
			v := p.v
			if dis[v][k] < p.d {
				continue
			}
			for i, e := range g[v] {
				w, wt := e.to, e.wt
				if e.del && k > 0 {
					if newD := dis[w][0] + diff - dis[v][1]; newD > wt {
						wt = newD
						g[v][i].wt = newD
						g[w][e.rid].wt = newD
					}
				}
				if newD := dis[v][k] + wt; newD < dis[w][k] {
					dis[w][k] = newD
					h.push(pair15{w, newD})
				}
			}
		}
	}
	dij(0)
	diff = l - dis[t][0]
	if diff < 0 {
		Fprint(out, "NO")
		return
	}
	dij(1)
	if dis[t][1] < l {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for v, es := range g {
		for _, e := range es {
			if e.to > v {
				Fprintln(out, v, e.to, e.wt)
			}
		}
	}
}

//func main() { CF715B(os.Stdin, os.Stdout) }
