package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, l, ans int
	Fscan(in, &n, &l)
	a := make([]int, n, n+1)
	for i := range a {
		Fscan(in, &a[i])
		l -= a[i]
	}
	if l > 0 {
		a = append(a, l)
	}
	h := hp{a}
	heap.Init(&h)
	for h.Len() > 1 {
		h.IntSlice[0] += heap.Pop(&h).(int)
		ans += h.IntSlice[0]
		heap.Fix(&h, 0)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
type hp struct{ sort.IntSlice }
func (hp) Push(interface{})    {}
func (h *hp) Pop() interface{} { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
