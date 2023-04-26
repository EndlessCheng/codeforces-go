package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF915F(_r io.Reader, out io.Writer) {
	buf := make([]byte, 1<<12)
	_i := len(buf)
	rc := func() byte {
		if _i >= len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	n := r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}
	es := make([]struct{ v, w int }, n-1)
	for i := range es {
		v, w := r()-1, r()-1
		if a[v] > a[w] {
			v, w = w, v
		}
		es[i].v = v
		es[i].w = w
	}

	sort.Slice(es, func(i, j int) bool { return a[es[i].w] < a[es[j].w] })
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	ans := int64(0)
	for _, e := range es {
		v, w := find(e.v), find(e.w)
		ans += int64(sz[v]) * int64(sz[w]) * int64(a[e.w])
		fa[v] = w
		sz[w] += sz[v]
	}

	sort.Slice(es, func(i, j int) bool { return a[es[i].v] > a[es[j].v] })
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	for _, e := range es {
		v, w := find(e.v), find(e.w)
		ans -= int64(sz[v]) * int64(sz[w]) * int64(a[e.v])
		fa[v] = w
		sz[w] += sz[v]
	}
	Fprint(out, ans)
}

//func main() { CF915F(os.Stdin, os.Stdout) }
