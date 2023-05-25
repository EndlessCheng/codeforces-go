package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1693C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	rg := make([][]int, n+1)
	deg := make([]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		rg[w] = append(rg[w], v)
		deg[v]++
	}
	dis := make([]int, n+1)
	for i := 1; i < n; i++ {
		dis[i] = 1e9
	}
	h := hp93{{n, 0}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair93)
		w := top.v
		if top.dis > dis[w] {
			continue
		}
		for _, v := range rg[w] {
			newD := top.dis + deg[v]
			if newD < dis[v] {
				dis[v] = newD
				heap.Push(&h, pair93{v, newD})
			}
			deg[v]--
		}
	}
	Fprint(out, dis[1])
}

//func main() { CF1693C(os.Stdin, os.Stdout) }
type pair93 struct{ v, dis int }
type hp93 []pair93
func (h hp93) Len() int           { return len(h) }
func (h hp93) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp93) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp93) Push(v any)        { *h = append(*h, v.(pair93)) }
func (h *hp93) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
