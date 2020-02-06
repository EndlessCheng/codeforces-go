package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

type intHeap struct{ sort.IntSlice }

func (h *intHeap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *intHeap) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

// github.com/EndlessCheng/codeforces-go
func p4447(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	endsWith := map[int]*intHeap{}
	for _, v := range a {
		minLen := 0
		if h, ok := endsWith[v-1]; ok {
			minLen = Pop(h).(int)
			if h.Len() == 0 {
				delete(endsWith, v-1)
			}
		}
		if _, ok := endsWith[v]; !ok {
			endsWith[v] = &intHeap{}
		}
		Push(endsWith[v], minLen+1)
	}
	ans := int(1e9)
	for _, h := range endsWith {
		if h.IntSlice[0] < ans {
			ans = h.IntSlice[0]
		}
	}
	Fprint(out, ans)
}

func main() { p4447(os.Stdin, os.Stdout) }
