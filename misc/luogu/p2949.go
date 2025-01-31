package main

import (
	"container/heap"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func p2949(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	type pair struct{ d, p int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].d, &a[i].p)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.d - b.d })
	h := &hp2949{}
	for _, p := range a {
		ans += p.p
		heap.Push(h, p.p)
		if h.Len() > p.d {
			ans -= heap.Pop(h).(int)
		}
	}
	Fprint(out, ans)
}

//func main() { p2949(bufio.NewReader(os.Stdin), os.Stdout) }

type hp2949 struct{ sort.IntSlice }
func (h *hp2949) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2949) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
