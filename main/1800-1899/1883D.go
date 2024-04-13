package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1883D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, l, r int
	var op string
	minR, maxL := hp83{nil, map[int]int{}}, hp83{nil, map[int]int{}}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == "+" {
			minR.push(r)
			maxL.push(-l)
		} else {
			minR.del(r)
			maxL.del(-l)
		}
		minR.do()
		maxL.do()
		if minR.Len() > 0 && minR.IntSlice[0] < -maxL.IntSlice[0] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1883D(os.Stdin, os.Stdout) }

type hp83 struct {
	sort.IntSlice
	todo map[int]int
}
func (h *hp83) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp83) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp83) del(v int)  { h.todo[v]++ }
func (h *hp83) push(v int) {
	if h.todo[v] > 0 {
		h.todo[v]--
	} else {
		heap.Push(h, v)
	}
}
func (h *hp83) do() {
	for h.Len() > 0 && h.todo[h.IntSlice[0]] > 0 {
		h.todo[h.IntSlice[0]]--
		heap.Pop(h)
	}
}
