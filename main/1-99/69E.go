package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF69E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	cnt := map[int]int{}
	h := hp69{}
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		cnt[v]++
		if cnt[v] == 1 {
			heap.Push(&h, v)
		}
		if i >= k-1 {
			for h.Len() > 0 && cnt[h.IntSlice[0]] != 1 {
				heap.Pop(&h)
			}
			if h.Len() > 0 {
				Fprintln(out, h.IntSlice[0])
			} else {
				Fprintln(out, "Nothing")
			}
			v := a[i-k+1]
			cnt[v]--
			if cnt[v] == 1 {
				heap.Push(&h, v)
			}
		}
	}
}

//func main() { CF69E(os.Stdin, os.Stdout) }
type hp69 struct{ sort.IntSlice }
func (h hp69) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp69) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp69) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
