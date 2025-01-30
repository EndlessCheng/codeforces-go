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
	cc int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa, n}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) (isNewMerge bool) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u.fa[x] = y
	u.cc--
	return true
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }

func run(in io.Reader, out io.Writer) {
	var n, m, v, w, z int
	Fscan(in, &n, &m)
	u := newUnionFind(n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &z)
		u.merge(v-1, w-1)
	}
	Fprint(out, u.cc)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
