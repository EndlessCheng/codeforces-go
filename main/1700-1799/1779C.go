package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1779C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		f := func(a []int) {
			h := hp79{}
			s := int64(0)
			for _, v := range a {
				heap.Push(&h, v)
				s += int64(v)
				if s < 0 {
					ans++
					s -= int64(heap.Pop(&h).(int) * 2)
				}
			}
		}
		Fscan(in, &n, &m, &v)
		a := make([]int, m-1)
		for i := m - 2; i >= 0; i-- {
			Fscan(in, &a[i])
			a[i] = -a[i]
		}
		f(a)
		a = make([]int, n-m)
		for i := range a {
			Fscan(in, &a[i])
		}
		f(a)
		Fprintln(out, ans)
	}
}

//func main() { CF1779C(os.Stdin, os.Stdout) }
type hp79 struct{ sort.IntSlice }
func (h *hp79) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp79) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
