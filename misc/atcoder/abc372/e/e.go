package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
type uf struct {
	fa []int
	a  []sort.IntSlice
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	sz := make([]sort.IntSlice, n)
	for i := range fa {
		fa[i] = i
		sz[i] = []int{i}
	}
	return uf{fa, sz}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) (newRoot int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return -1
	}
	u.fa[x] = y
	u.a[y] = append(u.a[y], u.a[x]...)
	sort.Sort(sort.Reverse(u.a[y]))
	if len(u.a[y]) > 10 {
		u.a[y] = u.a[y][:10]
	}
	return y
}

func (u uf) val(x, k int) int {
	a := u.a[u.find(x)]
	if len(a) < k {
		return -1
	}
	return a[k-1]
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, v, w int
	Fscan(in, &n, &q)
	u := newUnionFind(n + 1)
	for ; q > 0; q-- {
		Fscan(in, &op, &v, &w)
		if op == 1 {
			u.merge(v, w)
		} else {
			Fprintln(out, u.val(v, w))
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
