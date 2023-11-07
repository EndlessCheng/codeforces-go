package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1140C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans, sum int
	Fscan(in, &n, &k)
	a := make([]struct{ len, b int }, n)
	for i := range a {
		Fscan(in, &a[i].len, &a[i].b)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].b > a[j].b })

	h := &hp40{}
	for _, p := range a {
		heap.Push(h, p.len)
		sum += p.len
		for h.Len() > k {
			sum -= heap.Pop(h).(int)
		}
		ans = max(ans, sum*p.b)
	}
	Fprint(out, ans)
}

//func main() { CF1140C(os.Stdin, os.Stdout) }
type hp40 struct{ sort.IntSlice }
func (h *hp40) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp40) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
