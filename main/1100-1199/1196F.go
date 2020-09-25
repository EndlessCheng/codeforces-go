package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type edge96 struct {
	v, w int
	d    int64
}
type hp96 []edge96

func (h hp96) Len() int            { return len(h) }
func (h hp96) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h hp96) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp96) Push(v interface{}) { *h = append(*h, v.(edge96)) }
func (h *hp96) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp96) push(v edge96)      { heap.Push(h, v) }
func (h *hp96) pop() edge96        { return heap.Pop(h).(edge96) }

func CF1196F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, w, wt int
	Fscan(in, &n, &m, &k)
	q := hp96{}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		q = append(q, edge96{v, w, int64(wt)})
	}
	heap.Init(&q)

	ans := int64(0)
	dis := make([]map[int]int64, n+1)
	for i := 1; i <= n; i++ {
		dis[i] = map[int]int64{}
	}
	for k > 0 && len(q) > 0 {
		e := q.pop()
		v, w := e.v, e.w
		if dis[v][w] > 0 {
			continue
		}
		for v, d := range dis[v] {
			q.push(edge96{w, v, d + e.d})
		}
		for w, d := range dis[w] {
			q.push(edge96{v, w, d + e.d})
		}
		dis[v][w] = e.d
		dis[w][v] = e.d
		ans = e.d
		k--
	}
	Fprint(out, ans)
}

//func main() { CF1196F(os.Stdin, os.Stdout) }
