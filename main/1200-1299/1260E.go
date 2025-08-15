package main

import (
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1260E(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	h := hp60{}
	for i := n - 1; ; i-- {
		if a[i] < 0 {
			Fprint(out, ans)
			return
		}
		heap.Push(&h, a[i])
		if (i+1)&i == 0 {
			ans += heap.Pop(&h).(int)
		}
	}
}

//func main() { cf1260E(bufio.NewReader(os.Stdin), os.Stdout) }

type hp60 struct{ sort.IntSlice }
func (h *hp60) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp60) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
