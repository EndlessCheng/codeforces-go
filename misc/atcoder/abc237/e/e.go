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

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1e18
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		v := p.v
		if dis[v] < p.d {
			continue
		}
		for _, w := range g[v] {
			if d := dis[v] + max(a[w]-a[v], 0); d < dis[w] {
				dis[w] = d
				heap.Push(&h, pair{w, d})
			}
		}
	}
	for i, d := range dis {
		ans = max(ans, a[0]-a[i]-d)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
