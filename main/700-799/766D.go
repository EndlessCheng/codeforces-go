package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF766D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, d int
	var x, y string
	Fscan(in, &n, &m, &q)
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			ffx := find(fa[x])
			dis[x] += dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	delta := func(x, y int) int { return ((dis[x]-dis[y])%2 + 2) % 2 }
	merge := func(from, to int, d int) bool {
		if fFrom, fTo := find(from), find(to); fFrom != fTo {
			dis[fFrom] = d + dis[to] - dis[from]
			fa[fFrom] = fTo
			return true
		}
		return delta(from, to) == d
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	id := make(map[string]int, n)
	for ; n > 0; n-- {
		Fscan(in, &x)
		id[x] = len(id)
	}
	for ; m > 0; m-- {
		Fscan(in, &d, &x, &y)
		if merge(id[x], id[y], d-1) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &x, &y)
		if a, b := id[x], id[y]; same(a, b) {
			Fprintln(out, delta(a, b)+1)
		} else {
			Fprintln(out, 3)
		}
	}
}

//func main() { CF766D(os.Stdin, os.Stdout) }