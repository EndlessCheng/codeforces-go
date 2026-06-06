package main

import (
	"container/heap"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf639D(in io.Reader, out io.Writer) {
	var n, k, b, c int
	Fscan(in, &n, &k, &b, &c)
	b = min(b, c*5)
	t := make([]int, n)
	for i := range t {
		Fscan(in, &t[i])
	}
	slices.Sort(t)

	ans := int(1e18)
	for i := range 5 {
		T := int(2e9) + i
		h := &hp39{}
		heap.Init(h)
		sum := 0
		for j := 0; j < n; j++ {
			nd := (T-t[j])/5*b + (T-t[j])%5*c
			sum += nd
			heap.Push(h, nd)
			if h.Len() > k {
				sum -= heap.Pop(h).(int)
			}
			if h.Len() == k {
				ans = min(ans, sum-(T-t[j])/5*b*k)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf639D(bufio.NewReader(os.Stdin), os.Stdout) }

type hp39 struct{ sort.IntSlice }
func (h hp39) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp39) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp39) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
