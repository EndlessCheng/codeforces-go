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
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
	}

	h := hp{[]int{1}}
	vis := make([]bool, n+1)
	vis[1] = true
	cnt := 0
	for k := 1; k <= n; k++ {
		for h.Len() > 0 && h.IntSlice[0] <= k {
			v := heap.Pop(&h).(int)
			cnt++
			for _, w := range g[v] {
				if !vis[w] {
					vis[w] = true
					heap.Push(&h, w)
				}
			}
		}
		if cnt < k {
			Fprintln(out, -1)
		} else {
			Fprintln(out, h.Len())
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
