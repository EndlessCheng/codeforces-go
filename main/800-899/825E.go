package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func CF825E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	deg := make([]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[w] = append(g[w], v)
		deg[v]++
	}

	h := hp{}
	for i, d := range deg {
		if d == 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(&h)
	ans := make([]interface{}, n)
	for l := n; l > 0; l-- {
		v := h.pop()
		ans[v] = l
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				h.push(w)
			}
		}
	}
	Fprintln(out, ans...)
}

//func main() { CF825E(os.Stdin, os.Stdout) }
