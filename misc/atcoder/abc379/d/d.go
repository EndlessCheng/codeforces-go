package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q, op, x, add int
	h := hp{}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op)
		if op == 1 {
			heap.Push(&h, -add)
		} else if op == 2 {
			Fscan(in, &x)
			add += x
		} else {
			Fscan(in, &x)
			cnt := 0
			for h.Len() > 0 && h.IntSlice[0]+add >= x {
				heap.Pop(&h)
				cnt++
			}
			Fprintln(out, cnt)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
