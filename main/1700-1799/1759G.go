package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1759G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pos := make([]int, n+1)
		for i := 1; i < n; i += 2 {
			Fscan(in, &a[i])
			pos[a[i]] = i
		}
		h := hp59{}
		for i := n; i > 0; i-- {
			if pos[i] > 0 {
				heap.Push(&h, pos[i])
			} else if h.Len() > 0 {
				a[heap.Pop(&h).(int)-1] = i
			} else {
				Fprintln(out, -1)
				continue o
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1759G(os.Stdin, os.Stdout) }
type hp59 struct{ sort.IntSlice }
func (h hp59) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp59) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp59) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
