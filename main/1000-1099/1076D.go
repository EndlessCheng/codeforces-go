package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pr76 struct {
	v int
	d int64
}
type hp76 []pr76

func (h hp76) Len() int              { return len(h) }
func (h hp76) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp76) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp76) Push(v interface{})   { *h = append(*h, v.(pr76)) }
func (h *hp76) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp76) push(v pr76)          { heap.Push(h, v) }
func (h *hp76) pop() pr76            { return heap.Pop(h).(pr76) }

func CF1076D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, up, v, w, wt int
	Fscan(in, &n, &m, &up)
	type neighbor struct{ to, wt, i int }
	g := make([][]neighbor, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], neighbor{w, wt, i})
		g[w] = append(g[w], neighbor{v, wt, i})
	}

	type pair struct{ f, i int }
	fa := make([]pair, n)
	const inf int64 = 1e18
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0
	q := hp76{{}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if dist[v] < p.d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + int64(e.wt); newD < dist[w] {
				dist[w] = newD
				fa[w] = pair{v, e.i} // 构建最短路树（有向，从树根指向父节点）
				q.push(pr76{w, newD})
			}
		}
	}

	// 按照 fa 的方向拓扑排序
	deg := make([]int, n)
	for _, p := range fa {
		deg[p.f]++
	}
	ans := []interface{}{}
	q2 := []int{}
	for i, d := range deg {
		if d == 0 {
			q2 = append(q2, i)
		}
	}
	for len(q2) > 0 {
		p := fa[q2[0]]
		q2 = q2[1:]
		ans = append(ans, p.i)
		if p.f == 0 {
			continue
		}
		w := p.f
		if deg[w]--; deg[w] == 0 {
			q2 = append(q2, w)
		}
	}
	if len(ans) > up {
		ans = ans[len(ans)-up:]
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { CF1076D(os.Stdin, os.Stdout) }
