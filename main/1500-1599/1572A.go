package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// O(n) 做法是如果 v 需要 w 就连 v->w，如果 w 在 v 后面（w>v），那么边权为 1，否则边权为 0，答案就是从 0 出发的最长路

// github.com/EndlessCheng/codeforces-go
type hp72 struct{ sort.IntSlice }
func (h *hp72) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp72) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func CF1572A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		d := make([]int, n)
		h := hp72{}
		for w := range d {
			Fscan(in, &d[w])
			for i := 0; i < d[w]; i++ {
				Fscan(in, &v)
				g[v-1] = append(g[v-1], w)
			}
			if d[w] == 0 {
				h.IntSlice = append(h.IntSlice, w)
			}
		}
		ans, c := 0, 0
		for nxt := []int{}; h.Len() > 0; nxt = nil {
			for h.Len() > 0 {
				c++
				v := heap.Pop(&h).(int)
				for _, w := range g[v] {
					if d[w]--; d[w] == 0 {
						if w > v {
							heap.Push(&h, w)
						} else {
							nxt = append(nxt, w)
						}
					}
				}
			}
			h.IntSlice = nxt
			heap.Init(&h)
			ans++
		}
		if c < n {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1572A(os.Stdin, os.Stdout) }
