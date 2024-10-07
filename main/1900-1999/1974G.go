package main

import (
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1974G(in io.Reader, out io.Writer) {
	var T, n, x, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		ans, s := n, 0
		h := &hp74{}
		for ; n > 0; n-- {
			Fscan(in, &c)
			s -= c
			heap.Push(h, c)
			for s < 0 {
				s += heap.Pop(h).(int)
				ans--
			}
			s += x
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1974G(bufio.NewReader(os.Stdin), os.Stdout) }
type hp74 struct{ sort.IntSlice }
func (h hp74) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp74) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp74) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
