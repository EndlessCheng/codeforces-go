package main

import (
	"container/heap"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func p11457(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ end, t int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].end, &a[i].t)
			a[i].end += a[i].t
		}
		// 按照结束时间排序
		slices.SortFunc(a, func(a, b pair) int { return a.end - b.end })
		ans, sumT := 0, 0
		h := &hp11457{}
		for _, p := range a {
			ans++
			sumT += p.t
			heap.Push(h, p.t)
			if sumT > p.end {
				ans--
				sumT -= heap.Pop(h).(int)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { p11457(bufio.NewReader(os.Stdin), os.Stdout) }

type hp11457 struct{ sort.IntSlice }
func (h hp11457) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp11457) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp11457) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
