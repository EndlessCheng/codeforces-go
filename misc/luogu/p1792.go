package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type node1792 struct{ l, r, v int }
var a1792 []node1792

type hp1792 struct{ sort.IntSlice }
func (h hp1792) Less(i, j int) bool  { return a1792[h.IntSlice[i]].v > a1792[h.IntSlice[j]].v }
func (h *hp1792) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp1792) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func p1792(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, ans int
	Fscan(in, &n, &k)
	if k > n/2 {
		Fprint(out, "Error!")
		return
	}
	a1792 = make([]node1792, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a1792[i] = node1792{(i - 1 + n) % n, (i + 1) % n, v}
	}

	del := func(i int) { a1792[a1792[i].l].r, a1792[a1792[i].r].l = a1792[i].r, a1792[i].l }

	vis := make([]bool, n)
	h := &hp1792{make([]int, n)}
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
		ans += a1792[i].v
		l, r := a1792[i].l, a1792[i].r
		vis[l] = true
		vis[r] = true
		a1792[i].v = a1792[l].v + a1792[r].v - a1792[i].v
		heap.Fix(h, 0)
		del(l)
		del(r)
		k--
	}
	Fprint(out, ans)
}

//func main() { p1792(os.Stdin, os.Stdout) }
