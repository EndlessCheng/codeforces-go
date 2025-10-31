package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func p1792(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, ans int
	Fscan(in, &n, &k)
	if k > n/2 {
		Fprint(out, "Error!")
		return
	}
	a792 = make([]node792, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a792[i] = node792{(i - 1 + n) % n, (i + 1) % n, v}
	}

	del := func(i int) { a792[a792[i].l].r, a792[a792[i].r].l = a792[i].r, a792[i].l }

	vis := make([]bool, n)
	h := &hp792{make([]int, n)}
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
		ans += a792[i].v
		l, r := a792[i].l, a792[i].r
		vis[l] = true
		vis[r] = true
		a792[i].v = a792[l].v + a792[r].v - a792[i].v
		heap.Fix(h, 0)
		del(l)
		del(r)
		k--
	}
	Fprint(out, ans)
}

//func main() { p1792(os.Stdin, os.Stdout) }

type node792 struct{ l, r, v int }
var a792 []node792
type hp792 struct{ sort.IntSlice }
func (h hp792) Less(i, j int) bool  { return a792[h.IntSlice[i]].v > a792[h.IntSlice[j]].v }
func (h *hp792) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp792) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
