package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair86 struct{ v, d int }
type hp86 []pair86

func (h hp86) Len() int              { return len(h) }
func (h hp86) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp86) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp86) Push(v interface{})   { *h = append(*h, v.(pair86)) }
func (h *hp86) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp86) push(v pair86)        { heap.Push(h, v) }
func (h *hp86) pop() pair86          { return heap.Pop(h).(pair86) }

func CF1486E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9
	type nb struct{ to, wt int }

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	g := make([][]nb, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	dist := make([]int, n)
	mi := make([]int, n)
	for i := range dist {
		dist[i] = inf
		mi[i] = inf
	}
	dist[0] = 0
	q := hp86{{}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if p.d > dist[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			// 优先队列出来的点肯定是先近后远
			// 假设 v 比 v' 先出来
			// 对于任意中间点 w
			// 如果 v 可以用距离 d 和 w 做中转，那么没有必要再考虑 v' 用 >=d 的距离和 w 做中转了
			if e.wt >= mi[w] {
				continue
			}
			mi[w] = e.wt
			for _, e2 := range g[w] {
				u := e2.to
				if newD := dist[v] + (e.wt+e2.wt)*(e.wt+e2.wt); newD < dist[u] {
					dist[u] = newD
					q.push(pair86{u, newD})
				}
			}
		}
	}
	for _, v := range dist {
		if v == inf {
			v = -1
		}
		Fprint(out, v, " ")
	}
}

//func main() { CF1486E(os.Stdin, os.Stdout) }
