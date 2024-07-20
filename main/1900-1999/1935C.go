package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"slices"
	"sort"
)

func cf1935C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, lim int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &lim)
		type pair struct{ a, b int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].a, &a[i].b)
		}
		slices.SortFunc(a, func(p, q pair) int { return p.b - q.b })

		ans := 0
		for i, p := range a {
			if p.a <= lim {
				ans = max(ans, 1)
			}
			h := hp35{}
			s := 0
			for j := i - 1; j >= 0 && p.a+p.b-a[j].b < lim; j-- {
				q := a[j]
				heap.Push(&h, q.a)
				s += q.a
				for h.Len() > 0 && p.a+p.b-q.b+s > lim {
					s -= heap.Pop(&h).(int)
				}
				ans = max(ans, 1+h.Len())
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1935C(bufio.NewReader(os.Stdin), os.Stdout) }
type hp35 struct{ sort.IntSlice }
func (h hp35) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp35) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp35) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
