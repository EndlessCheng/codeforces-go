package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type uf struct {
	fa []int
	sz []int
	c  []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	sz := make([]int, n)
	c := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
		c[i] = i
	}
	return uf{fa, sz, c}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.sz[y] += u.sz[x]
}

func (u uf) size(x int) int {
	return u.sz[u.find(x)]
}

func (u uf) color(x int) int {
	return u.c[u.find(x)]
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, x, c int
	Fscan(in, &n, &q)
	u := newUnionFind(n + 1)
	color := make([]int, n+1)
	for i := range color {
		color[i] = 1
	}

	for range q {
		Fscan(in, &op, &x)
		if op == 2 {
			Fprintln(out, color[x])
			continue
		}

		Fscan(in, &c)
		x = u.find(x)
		color[u.c[x]] -= u.sz[x]
		u.c[x] = c
		color[u.c[x]] += u.sz[x]

		l := x - u.sz[x]
		if l > 0 && u.c[l] == c {
			u.merge(l, x)
		}
		if x < n && u.color(x+1) == c {
			u.merge(x, x+1)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
