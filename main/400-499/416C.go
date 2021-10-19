package main

import (
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
var a16 [1e3]struct{ v, p, i int }
type hp16 struct{ sort.IntSlice }
func (h hp16) Less(i, j int) bool  { return a16[h.IntSlice[i]].p < a16[h.IntSlice[j]].p }
func (h *hp16) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp16) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func CF416C(in io.Reader, out io.Writer) {
	var n, m, sum int
	Fscan(in, &n)
	a := a16[:n]
	for i := range a {
		Fscan(in, &a[i].v, &a[i].p)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })
	Fscan(in, &m)
	type pair struct{ v, i int }
	t := make([]pair, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &t[i].v)
		t[i].i = i
	}
	sort.Slice(t, func(i, j int) bool { return t[i].v > t[j].v })

	h, j := hp16{}, 0
	for i, p := range a {
		for ; t[j].v >= p.v; j++ {
		}
		heap.Push(&h, i)
		if h.Len() > j {
			heap.Pop(&h)
		}
	}
	b := h.IntSlice
	for _, i := range b {
		sum += a[i].p
	}
	Fprintln(out, len(b), sum)
	sort.Slice(b, func(i, j int) bool { return a[b[i]].v > a[b[j]].v })
	j = -1
	for _, i := range b {
		for j++; t[j].v < a[i].v; j++ {
		}
		Fprintln(out, a[i].i, t[j].i)
	}
}

//func main() { CF416C(os.Stdin, os.Stdout) }
