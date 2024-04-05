package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1732D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, x int
	var op string
	Fscan(in, &q)
	mex := map[int]int{}
	has := map[int]bool{}
	belong := map[int][]*lazyHeap32{}
	missHs := map[int]*lazyHeap32{}
	for ; q > 0; q-- {
		Fscan(in, &op, &x)
		if op == "+" {
			has[x] = true
			for _, miss := range belong[x] {
				miss.toDel[x] = true
			}
		} else if op == "-" {
			delete(has, x)
			for _, miss := range belong[x] {
				if miss.toDel[x] {
					delete(miss.toDel, x)
				} else {
					heap.Push(miss, x)
				}
			}
		} else {
			if missHs[x] == nil {
				missHs[x] = &lazyHeap32{[]int{}, map[int]bool{}}
			}
			miss := missHs[x]
			for miss.Len() > 0 && miss.toDel[miss.IntSlice[0]] {
				delete(miss.toDel, miss.IntSlice[0])
				heap.Pop(miss)
			}
			if miss.Len() > 0 {
				Fprintln(out, miss.IntSlice[0])
				continue
			}
			if mex[x] == 0 {
				mex[x] = x
			}
			for has[mex[x]] {
				belong[mex[x]] = append(belong[mex[x]], miss)
				mex[x] += x
			}
			Fprintln(out, mex[x])
		}
	}
}

//func main() { cf1732D2(os.Stdin, os.Stdout) }

type lazyHeap32 struct {
	sort.IntSlice
	toDel map[int]bool
}

func (h *lazyHeap32) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap32) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
