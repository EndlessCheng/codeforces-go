package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF482B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ l, r int }

	var n, q, l, r, v int
	Fscan(in, &n, &q)
	fas := [30][]int{}
	for i := range fas {
		fas[i] = make([]int, n+1)
		for j := range fas[i] {
			fas[i][j] = j
		}
	}
	var find func([]int, int) int
	find = func(fa []int, x int) int {
		if fa[x] != x {
			fa[x] = find(fa, fa[x])
		}
		return fa[x]
	}
	same := func(fa []int, x, y int) bool { return find(fa, x) == find(fa, y) }
	mergeRange := func(fa []int, l, r int) {
		for i := find(fa, l); i < r; i = find(fa, i+1) {
			fa[i] = r
		}
	}
	check0 := [30][]pair{}
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &v)
		l--
		for i, fa := range fas {
			if v>>i&1 == 0 {
				check0[i] = append(check0[i], pair{l, r})
			} else {
				mergeRange(fa, l, r)
			}
		}
	}
	for i, qs := range check0 {
		fa := fas[i]
		for _, q := range qs {
			if same(fa, q.l, q.r) {
				Fprintln(out, "NO")
				return
			}
		}
	}
	ans := make([]int, n)
	for j, fa := range fas {
		for i, f := range fa {
			if f != i {
				ans[i] |= 1 << j
			}
		}
	}
	Fprintln(out, "YES")
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF482B(os.Stdin, os.Stdout) }
