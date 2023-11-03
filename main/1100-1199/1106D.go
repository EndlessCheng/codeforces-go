package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1106D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	h := hp6{[]int{1}}
	vis := make([]bool, n+1)
	vis[1] = true
	for h.Len() > 0 {
		v := heap.Pop(&h).(int)
		Fprint(out, v, " ")
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				heap.Push(&h, w)
			}
		}
	}
}

//func main() { CF1106D(os.Stdin, os.Stdout) }
type hp6 struct{ sort.IntSlice }
func (h *hp6) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp6) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
