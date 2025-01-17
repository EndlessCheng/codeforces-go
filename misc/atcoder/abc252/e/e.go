package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt, i int }
	g := make([][]nb, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt, i})
		g[w] = append(g[w], nb{v, wt, i})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1e18
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair)
		v := top.v
		d := top.dis
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

	vis := make([]bool, n)
	for v, ws := range g {
		for _, e := range ws {
			w := e.to
			if vis[w] {
				continue
			}
			if dis[v]+e.wt == dis[w] {
				Fprint(out, e.i, " ")
				vis[w] = true
			}
		}
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
