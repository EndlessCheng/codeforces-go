package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

type pri struct {
	r, i   int
	belong *phi
}
type hp []*pri

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].r > h[j].r }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].i = i; h[j].i = j }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(*pri)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v *pri)        { v.i = len(*h); heap.Push(h, v) }
func (h *hp) remove(i int)       { heap.Remove(h, i) }
func (h hp) empty() bool         { return len(h) == 0 }
func (h hp) top() int            { return h[0].r }

type phi struct {
	h hp
	i int
}
type hhp []*phi

func (h hhp) Len() int            { return len(h) }
func (h hhp) Less(i, j int) bool  { return h[i].h.top() < h[j].h.top() }
func (h hhp) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].i = i; h[j].i = j }
func (h *hhp) Push(v interface{}) { *h = append(*h, v.(*phi)) }
func (h *hhp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hhp) push(v *phi)        { v.i = len(*h); heap.Push(h, v) }
func (h *hhp) remove(i int)       { heap.Remove(h, i) }
func (h *hhp) fix(i int)          { heap.Fix(h, i) }
func (h hhp) top() hp             { return h[0].h }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	his := [2e5 + 1]phi{}
	var n, q, p, id int
	Fscan(in, &n, &q)
	a := make([]pri, n)
	for i := range a {
		Fscan(in, &a[i].r, &p)
		a[i].i = len(his[p].h)
		his[p].h = append(his[p].h, &a[i])
	}

	hh := hhp{}
	for i := range his {
		if hi := &his[i]; !hi.h.empty() {
			heap.Init(&hi.h)
			hi.i = id
			id++
			for _, ri := range hi.h {
				ri.belong = hi
			}
			hh = append(hh, hi)
		}
	}
	heap.Init(&hh)

	for ; q > 0; q-- {
		Fscan(in, &id, &p)
		ri := &a[id-1]
		hi := ri.belong
		hi.h.remove(ri.i)
		if hi.h.empty() {
			hh.remove(hi.i)
		} else {
			hh.fix(hi.i)
		}

		hi = &his[p]
		isNew := hi.h.empty()
		hi.h.push(ri)
		ri.belong = hi
		if isNew {
			hh.push(hi)
		} else {
			hh.fix(hi.i)
		}
		Fprintln(out, hh.top().top())
	}
}

func main() { run(os.Stdin, os.Stdout) }
