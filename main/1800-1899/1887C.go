package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1887C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, l, r, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		sd := make([]int, n+1)
		d := make([]int, n+1)
		h := hp87{}
		Fscan(in, &q)
		for range q {
			Fscan(in, &l, &r, &x)
			if x == 0 {
				continue
			}
			l--
			if d[l] == 0 {
				heap.Push(&h, l)
			}
			d[l] += x
			if d[r] == 0 {
				heap.Push(&h, r)
			}
			d[r] -= x
			for h.Len() > 0 && d[h.IntSlice[0]] == 0 {
				heap.Pop(&h)
			}
			if h.Len() > 0 && d[h.IntSlice[0]] < 0 {
				for _, i := range h.IntSlice {
					sd[i] += d[i]
					d[i] = 0
				}
				h.IntSlice = h.IntSlice[:0]
			}
		}

		s := 0
		for i, v := range a {
			s += sd[i]
			Fprint(out, v+s, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1887C(bufio.NewReader(os.Stdin), os.Stdout) }
type hp87 struct{ sort.IntSlice }
func (h *hp87) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp87) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
