package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type vd22 struct{ v, d int }
type hp22 []vd22

func (h hp22) Len() int              { return len(h) }
func (h hp22) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp22) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp22) Push(v interface{})   { *h = append(*h, v.(vd22)) }
func (h *hp22) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp22) push(v vd22)          { heap.Push(h, v) }
func (h *hp22) pop() vd22            { return heap.Pop(h).(vd22) }

func CF1422D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, sx, sy, tx, ty, x, y int
	Fscan(in, &n, &n, &sx, &sy, &tx, &ty)
	dist := make([]int, n)
	end := make([]int, n)
	q := make(hp22, n)
	type pt struct{ x, y, i int }
	ps := make([]pt, n)
	qs := make([]pt, n)
	for i := range dist {
		Fscan(in, &x, &y)
		dist[i] = min(abs(sx-x), abs(sy-y))
		end[i] = abs(x-tx) + abs(y-ty)
		q[i] = vd22{i, dist[i]}
		ps[i] = pt{x, y, i}
		qs[i] = pt{y, x, i}
	}

	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for _, ps := range [][]pt{ps, qs} {
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
		for i := 1; i < n; i++ {
			p, q := ps[i-1], ps[i]
			v, w, d := p.i, q.i, min(q.x-p.x, abs(q.y-p.y))
			g[v] = append(g[v], nb{w, d})
			g[w] = append(g[w], nb{v, d})
		}
	}

	heap.Init(&q)
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if dist[v] < p.d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				q.push(vd22{w, newD})
			}
		}
	}
	ans := abs(sx-tx) + abs(sy-ty)
	for i, d := range dist {
		ans = min(ans, d+end[i])
	}
	Fprint(out, ans)
}

//func main() { CF1422D(os.Stdin, os.Stdout) }
