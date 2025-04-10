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
			Fscan(in, &x)
			heap.Push(&h, x-add)
		} else if op == 2 {
			Fscan(in, &x)
			add += x
		} else {
			Fprintln(out, heap.Pop(&h).(int)+add)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
