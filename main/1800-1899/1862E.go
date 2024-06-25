package main

import (
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

func cf1862E(in io.Reader, out io.Writer) {
	var T, n, m, d, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &d)
		ans, s := 0, 0
		h := hp62{}
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if v <= 0 {
				continue
			}
			if h.Len() < m {
				s += v
				heap.Push(&h, v)
			} else if v > h.IntSlice[0] {
				s += v - h.replace(v)
			}
			ans = max(ans, s-i*d)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1862E(bufio.NewReader(os.Stdin), os.Stdout) }
type hp62 struct{ sort.IntSlice }
func (h *hp62) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp62) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp62) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return top }
