package main

import (
	"container/heap"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type uf99 struct {
	fa []int
	sz []int
}

func newUnionFind99(n int) uf99 {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return uf99{fa, sz}
}

func (u uf99) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf99) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.sz[y] += u.sz[x]
}

func cf899E(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)

	prev := make([]int, n)
	next := make([]int, n)
	for i := range n {
		prev[i] = i - 1
		next[i] = i + 1
	}
	del := func(i int) {
		l, r := prev[i], next[i]
		if l >= 0 {
			next[l] = r
		}
		if r < n {
			prev[r] = l
		}
	}

	u := newUnionFind99(n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if i > 0 && a[i] == a[i-1] {
			u.merge(i-1, i)
			del(i - 1)
		}
	}

	h := hp99{}
	for i, rt := range u.fa {
		if rt == i {
			heap.Push(&h, pair99{u.sz[i], i})
		}
	}

	for len(h) > 0 {
		p := heap.Pop(&h).(pair99)
		i := p.i
		if i != u.find(i) || p.sz != u.sz[i] {
			continue
		}
		l, r := prev[i], next[i]
		if l >= 0 && r < n && a[l] == a[r] {
			u.merge(l, r)
			j := u.find(r)
			heap.Push(&h, pair99{u.sz[j], j})
			del(l)
		}
		del(i)
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
