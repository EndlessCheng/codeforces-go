package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type uf60 []int

func newUnionFind60(n int) uf60 {
	fa := make(uf60, n)
	for i := range fa {
		fa[i] = i
	}
	return fa
}

func (u uf60) find(x int) int {
	if u[x] != x {
		u[x] = u.find(u[x])
	}
	return u[x]
}

func (u uf60) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u[x] = y
	return true
}

func (u uf60) same(x, y int) bool { return u.find(x) == u.find(y) }

func cf2060E(in io.Reader, out io.Writer) {
	var T, n, m1, m2 int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m1, &m2)
		type edge struct{ v, w int }
		es1 := make([]edge, m1)
		for i := range es1 {
			Fscan(in, &es1[i].v, &es1[i].w)
		}
		es2 := make([]edge, m2)
		uf2 := newUnionFind60(n + 1)
		for i := range es2 {
			Fscan(in, &es2[i].v, &es2[i].w)
			uf2.merge(es2[i].v, es2[i].w)
		}

		ans := 0
		uf1 := newUnionFind60(n + 1)
		for _, e := range es1 {
			if uf2.same(e.v, e.w) {
				uf1.merge(e.v, e.w)
			} else {
				ans++
			}
		}
		for _, e := range es2 {
			if uf1.merge(e.v, e.w) {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2060E(bufio.NewReader(os.Stdin), os.Stdout) }
