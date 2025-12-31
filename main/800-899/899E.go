package main

import (
	"container/heap"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf899E(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)

	pre := make([]int, n)
	nxt := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		pre[i] = i - 1
		nxt[i] = i + 1
		sz[i] = 1
	}
	del := func(i int) {
		l, r := pre[i], nxt[i]
		if l >= 0 {
			nxt[l] = r
		}
		if r < n {
			pre[r] = l
		}
	}
	merge := func(from, to int) {
		sz[to] += sz[from]
		sz[from] = 0
		del(from)
	}

	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if i > 0 && a[i] == a[i-1] {
			merge(i-1, i)
		}
	}

	h := hp99{}
	for i, s := range sz {
		if s > 0 {
			heap.Push(&h, pair99{s, i})
		}
	}

	for len(h) > 0 {
		p := heap.Pop(&h).(pair99)
		i := p.i
		if sz[i] != p.sz {
			continue
		}
		sz[i] = 0
		l, r := pre[i], nxt[i]
		del(i)
		if l >= 0 && r < n && a[l] == a[r] {
			merge(l, r)
			heap.Push(&h, pair99{sz[r], r})
		}
		ans++
	}
	Fprint(out, ans)
}

//func main() { cf899E(bufio.NewReader(os.Stdin), os.Stdout) }

type pair99 struct{ sz, i int }
type hp99 []pair99
func (h hp99) Len() int           { return len(h) }
func (h hp99) Less(i, j int) bool { return h[i].sz > h[j].sz || h[i].sz == h[j].sz && h[i].i < h[j].i }
func (h hp99) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp99) Push(v any)        { *h = append(*h, v.(pair99)) }
func (h *hp99) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
