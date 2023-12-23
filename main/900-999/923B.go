package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf923B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, dec, s int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	h := hp23{}
	for _, v := range a {
		Fscan(in, &dec)
		heap.Push(&h, v+s)
		res := 0
		for h.Len() > 0 && h.IntSlice[0] <= s+dec {
			res += heap.Pop(&h).(int) - s
		}
		res += h.Len() * dec
		Fprintln(out, res)
		s += dec
	}
}

//func main() { cf923B(os.Stdin, os.Stdout) }
type hp23 struct{ sort.IntSlice }
func (h *hp23) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp23) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
