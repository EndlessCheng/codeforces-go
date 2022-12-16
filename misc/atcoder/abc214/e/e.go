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
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ l, r int }, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
		}
		sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })
		h := &hp{}
		for i, left := 0, 1; i < n || h.Len() > 0; {
			for ; i < n && a[i].l == left; i++ {
				heap.Push(h, a[i].r)
			}
			if h.Len() == 0 {
				left = a[i].l
				continue
			}
			if heap.Pop(h).(int) < left {
				Fprintln(out, "No")
				continue o
			}
			left++
		}
		Fprintln(out, "Yes")
	}
}

func main() { run(os.Stdin, os.Stdout) }
type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
