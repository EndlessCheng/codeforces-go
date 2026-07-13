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
	var q, op, x int
	Fscan(in, &q)
	a := []int{}
	h := &hp{}
	for range q {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &x)
			a = append(a, x)
		} else if op == 3 {
			for _, v := range a {
				heap.Push(h, v)
			}
			a = a[:0]
		} else if h.Len() > 0 {
			Fprintln(out, heap.Pop(h))
		} else {
			Fprintln(out, a[0])
			a = a[1:]
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
