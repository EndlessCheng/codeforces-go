package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type edge struct{ to, wt int }
	g := make([][]edge, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}

	const inf int = 1e18
	d := make([]int, n+1)
	for i := range d {
		d[i] = inf
	}
	d[1] = 0
	from := make([]int, n+1)
	h := hp{{1, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		v := p.v
		if d[v] < p.dis {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := p.dis + e.wt; newD < d[w] {
				d[w] = newD
				from[w] = v
				heap.Push(&h, pair{w, newD})
			}
		}
	}

	if d[n] == inf {
		Fprint(out, -1)
		return
	}

	path := []int{}
	for x := n; x > 0; x = from[x] {
		path = append(path, x)
	}
	for i := len(path) - 1; i >= 0; i-- {
		Fprint(out, path[i], " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }

type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
