package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type node struct{ l, r, v int }
var a []node
func del(i int) { a[a[i].l].r, a[a[i].r].l = a[i].r, a[i].l }

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]].v > a[h.IntSlice[j]].v }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func p1792(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, ans int
	Fscan(in, &n, &k)
	if k > n/2 {
		Fprint(out, "Error!")
		return
	}
	a = make([]node, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a[i] = node{(i - 1 + n) % n, (i + 1) % n, v}
	}
	vis := make([]bool, n)
	h := &hp{make([]int, n)}
	for i := 0; i < n; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	for k > 0 {
		i := h.IntSlice[0]
		if vis[i] {
			heap.Pop(h)
			continue
		}
		ans += a[i].v
		l, r := a[i].l, a[i].r
		vis[l] = true
		vis[r] = true
		a[i].v = a[l].v + a[r].v - a[i].v
		heap.Fix(h, 0)
		del(l)
		del(r)
		k--
	}
	Fprint(out, ans)
}

//func main() { p1792(os.Stdin, os.Stdout) }
