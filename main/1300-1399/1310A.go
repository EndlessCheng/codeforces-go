package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1310A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]struct{ v, t int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
	}
	for i := range a {
		Fscan(in, &a[i].t)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })

	ans := int64(0)
	h := h1{}
	for i, cur, costSum := 0, 0, int64(0); i < n || h.Len() > 0; cur++ {
		if h.Len() == 0 {
			cur = a[i].v
		}
		for ; i < n && a[i].v == cur; i++ {
			heap.Push(&h, a[i].t)
			costSum += int64(a[i].t)
		}
		costSum -= int64(heap.Pop(&h).(int))
		ans += costSum
	}
	Fprint(out, ans)
}

//func main() { CF1310A(os.Stdin, os.Stdout) }

type h1 struct{ sort.IntSlice }

func (h h1) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *h1) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *h1) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
