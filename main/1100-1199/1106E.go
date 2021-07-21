package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type dw6 struct{ d, w int }
type pair6 struct {
	dw6
	i int
}
type mh6 []*pair6

func (h mh6) Len() int            { return len(h) }
func (h mh6) Less(i, j int) bool  { a, b := h[i], h[j]; return a.w > b.w || a.w == b.w && a.d > b.d }
func (h mh6) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].i = i; h[j].i = j }
func (h *mh6) Push(v interface{}) { *h = append(*h, v.(*pair6)) }
func (h *mh6) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh6) push(e dw6) *pair6  { p := &pair6{e, len(*h)}; heap.Push(h, p); return p }

func CF1106E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var n, m, k, s, t, d, w int
	Fscan(in, &n, &m, &k)
	es := make([][]dw6, n+1)
	for ; k > 0; k-- {
		Fscan(in, &s, &t, &d, &w)
		es[s] = append(es[s], dw6{d, w})
		if t < n {
			es[t+1] = append(es[t+1], dw6{-d, w})
		}
	}

	next := make([]dw6, n+1)
	h := mh6{}
	ptr := map[dw6][]*pair6{}
	for i, es := range es {
		for _, e := range es {
			if e.d > 0 {
				ptr[e] = append(ptr[e], h.push(e))
			} else {
				e.d = -e.d
				heap.Remove(&h, ptr[e][0].i)
				ptr[e] = ptr[e][1:]
			}
		}
		if len(h) > 0 {
			next[i] = h[0].dw6
		} else {
			next[i].d = i
		}
	}

	ans := int64(1e18)
	cur := make([]int64, n+2)
	for i := range cur {
		cur[i] = 1e18
	}
	cur[1] = 0
	for k := 0; k <= m; k++ {
		nxt := make([]int64, n+2)
		for i := range nxt {
			nxt[i] = 1e18
		}
		for i := 1; i <= n; i++ {
			nxt[i+1] = min(nxt[i+1], cur[i])
			cur[next[i].d+1] = min(cur[next[i].d+1], cur[i]+int64(next[i].w))
		}
		ans = min(ans, cur[n+1])
		cur = nxt
	}
	Fprint(out, ans)
}

//func main() { CF1106E(os.Stdin, os.Stdout) }
