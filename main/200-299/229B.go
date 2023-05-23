package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func CF229B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
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
	type seg struct{ l, r int }
	t := make([][]seg, n)
	for i := 0; i < n-1; i++ {
		pre := -2
		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &v)
			if v > pre+1 {
				t[i] = append(t[i], seg{v, v})
			} else {
				t[i][len(t[i])-1].r = v
			}
			pre = v
		}
	}
	t[n-1] = nil

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	if len(t[0]) > 0 && t[0][0].l == 0 {
		dis[0] = t[0][0].r + 1
	}
	h := hp29{{0, dis[0]}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair29)
		v := top.v
		if top.dis > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w, wt := e.to, e.wt
			newD := dis[v] + wt
			t := t[w]
			i := sort.Search(len(t), func(i int) bool { return t[i].r >= newD })
			if i < len(t) && t[i].l <= newD {
				newD = t[i].r + 1
			}
			if newD < dis[w] {
				dis[w] = newD
				heap.Push(&h, pair29{w, newD})
			}
		}
	}
	if dis[n-1] == math.MaxInt {
		Fprint(out, -1)
	} else {
		Fprint(out, dis[n-1])
	}
}

//func main() { CF229B(os.Stdin, os.Stdout) }
type pair29 struct{ v, dis int }
type hp29 []pair29
func (h hp29) Len() int              { return len(h) }
func (h hp29) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp29) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp29) Push(v interface{})   { *h = append(*h, v.(pair29)) }
func (h *hp29) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
