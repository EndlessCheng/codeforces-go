package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		g[v-1] = append(g[v-1], nb{w - 1, wt})
	}
	dis := make([]int, n)
	for st := range g {
		for i := range dis {
			dis[i] = 1e18
		}
		h := hp{{0, st}}
		for len(h) > 0 {
			p := heap.Pop(&h).(pair)
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
					heap.Push(&h, pair{newD, w})
				}
			}
		}
		d := dis[st]
		if d == 1e18 {
			d = -1
		}
		Fprintln(out, d)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

type pair struct{ dis, v int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
