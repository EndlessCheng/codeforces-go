package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair73 struct {
	v int
	d int64
}
type hp73 []pair73

func (h hp73) Len() int              { return len(h) }
func (h hp73) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp73) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp73) Push(v interface{})   { *h = append(*h, v.(pair73)) }
func (h *hp73) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp73) push(v pair73)        { heap.Push(h, v) }
func (h *hp73) pop() pair73          { return heap.Pop(h).(pair73) }

func CF1473E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	const inf int64 = 1e18
	dis := make([]int64, n*4)
	for i := range dis {
		dis[i] = inf
	}
	dis[0] = 0
	h := hp73{{}}
	for len(h) > 0 {
		pd := h.pop()
		x := pd.v
		if dis[x] < pd.d {
			continue
		}
		// 分成四层，分别表示原图、不算最大值、额外再算一遍最小值、不算最大值且额外再算一遍最小值
		// 从第一层到第四层就是要求的最短路
		// 即便当前的边不是最小或最大，后面转移的时候也会覆盖掉
		for _, e := range g[x>>2] {
			w, wt := e.to, int64(e.wt)
			y := w<<2 | x&3
			if newD := dis[x] + wt; newD < dis[y] {
				dis[y] = newD
				h.push(pair73{y, newD})
			}
			if newD := dis[x] + wt*2; x&1 == 0 && newD < dis[y|1] {
				dis[y|1] = newD
				h.push(pair73{y | 1, newD})
			}
			if newD := dis[x]; x>>1&1 == 0 && newD < dis[y|2] {
				dis[y|2] = newD
				h.push(pair73{y | 2, newD})
			}
		}
	}
	for i := 1; i < n; i++ {
		Fprint(out, min(dis[i<<2], dis[i<<2|3]), " ")
	}
}

//func main() { CF1473E(os.Stdin, os.Stdout) }
