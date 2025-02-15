package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, k, l int
	Fscan(in, &n, &m, &k, &l)
	color := make([]int, n)
	for i := range color {
		Fscan(in, &color[i])
	}
	const inf int = 1e18
	dis := make([]struct{ d, d2, c int }, n)
	for i := range dis {
		dis[i].d = inf
		dis[i].d2 = inf
	}
	h := hp{}
	for ; l > 0; l-- {
		var v int
		Fscan(in, &v)
		v--
		h = append(h, vd{0, v, color[v]})
	}
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	for len(h) > 0 {
		t := heap.Pop(&h).(vd)
		v, d, c := t.v, t.dis, t.c
		if dis[v].d == inf {
			dis[v].d = d
			dis[v].c = c
		} else if dis[v].d2 == inf && dis[v].c != c {
			dis[v].d2 = d
		} else {
			continue
		}
		for _, e := range g[v] {
			heap.Push(&h, vd{d + e.wt, e.to, c})
		}
	}
	for i, t := range dis {
		ans := -1
		if t.d < inf {
			if t.c != color[i] {
				ans = t.d
			} else if t.d2 < inf {
				ans = t.d2
			}
		}
		Fprint(out, ans, " ")
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

type vd struct{ dis, v, c int }
type hp []vd
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(vd)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
