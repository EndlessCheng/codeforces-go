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
func CF796C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := math.MaxInt
	h := lazyHeap{append(a[:0:0], a...), map[int]int{}}
	heap.Init(&h)
	for i, v := range a {
		mx := v
		h.delete(v)
		for _, w := range g[i] {
			mx = max(mx, a[w]+1)
			h.delete(a[w])
		}
		h.do()
		if h.Len() > 0 {
			mx = max(mx, h.IntSlice[0]+2)
		}
		ans = min(ans, mx)
		h.push(v)
		for _, w := range g[i] {
			h.push(a[w])
		}
	}
	Fprint(out, ans)
}

//func main() { CF796C(os.Stdin, os.Stdout) }

type lazyHeap struct {
	sort.IntSlice
	lazy map[int]int
}

func (h lazyHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *lazyHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *lazyHeap) push(v int) {
	if h.lazy[v] > 0 {
		h.lazy[v]--
	} else {
		heap.Push(h, v)
	}
}
func (h *lazyHeap) do() {
	for h.Len() > 0 && h.lazy[h.IntSlice[0]] > 0 {
		h.lazy[h.IntSlice[0]]--
		heap.Pop(h)
	}
}
func (h *lazyHeap) pop() int     { h.do(); return heap.Pop(h).(int) }
func (h *lazyHeap) top() int     { h.do(); return h.IntSlice[0] }
func (h *lazyHeap) delete(v int) { h.lazy[v]++ }
