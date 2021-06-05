package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type pair struct{ v, d int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type nb struct{ to, wt, i int }

	var n, m, k, v, w, wt int
	Fscan(in, &n, &m, &k)
	g := make([][]nb, n)
	for i := 0; i < m; i++ {
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
		vd := h.pop()
		v := vd.v
		if dis[v] < vd.d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dis[v] + e.wt; newD < dis[w] {
				dis[w] = newD
				h.push(pair{w, newD})
			}
		}
	}

	// 构建反向最短路树
	save := make([]bool, m)
	g2 := make([][]nb, n)
	deg := make([]int, n)
	vis := make([]bool, n)
	for v, es := range g {
		for _, e := range es {
			if w := e.to; !vis[w] && dis[v]+e.wt == dis[w] {
				g2[w] = append(g2[w], nb{v, 0, e.i})
				deg[v]++
				vis[w] = true
				save[e.i] = true
			}
		}
	}

	// 在反向最短路树上，按拓扑序删边
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for left := n - 1; left > k; {
		v := q[0]
		q = q[1:]
		for _, e := range g2[v] {
			save[e.i] = false
			if left--; left == k {
				break
			}
			w := e.to
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	ans := []interface{}{}
	for i, b := range save {
		if b {
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

func main() { run(os.Stdin, os.Stdout) }
