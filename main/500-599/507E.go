package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair7 struct{ d, ok int }
type hPair7 struct {
	v int
	pair7
}
type hp7 []hPair7

func (h hp7) Len() int              { return len(h) }
func (h hp7) Less(i, j int) bool    { a, b := h[i], h[j]; return a.d < b.d || a.d == b.d && a.ok > b.ok }
func (h hp7) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp7) Push(v interface{})   { *h = append(*h, v.(hPair7)) }
func (h *hp7) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp7) push(v hPair7)        { heap.Push(h, v) }
func (h *hp7) pop() hPair7          { return heap.Pop(h).(hPair7) }

func CF507E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, ok int
	Fscan(in, &n, &m)
	type nb struct{ to, ok, i int }
	g := make([][]nb, n)
	type edge struct{ v, w, ok int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w, &ok)
		v--
		w--
		es[i] = edge{v, w, ok}
		g[v] = append(g[v], nb{w, ok, i})
		g[w] = append(g[w], nb{v, ok, i})
	}

	dis := make([]pair7, n)
	for i := range dis {
		dis[i].d = 1e9
	}
	dis[0] = pair7{}
	vis := make([]bool, n)
	type vi struct{ v, i int }
	fa := make([]vi, n)
	for i := range fa {
		fa[i].v = -1
	}
	h := hp7{{}}
	for len(h) > 0 {
		vd := h.pop()
		v := vd.v
		if vis[v] {
			continue
		}
		vis[v] = true
		for _, e := range g[v] {
			w := e.to
			if newD, newOK := dis[v].d+1, dis[v].ok+e.ok; newD < dis[w].d || newD == dis[w].d && newOK > dis[w].ok {
				dis[w] = pair7{newD, newOK}
				fa[w] = vi{v, e.i}
				h.push(hPair7{w, dis[w]})
			}
		}
	}

	onPath := make([]int, m)
	for x := n - 1; fa[x].v >= 0; x = fa[x].v {
		onPath[fa[x].i] = 1
	}
	ans := []edge{}
	for i, e := range es {
		if e.ok != onPath[i] {
			ans = append(ans, edge{e.v + 1, e.w + 1, onPath[i]})
		}
	}
	Fprintln(out, len(ans))
	for _, e := range ans {
		Fprintln(out, e.v, e.w, e.ok)
	}
}

//func main() { CF507E(os.Stdin, os.Stdout) }
