package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
type uf []int

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return fa
}

func (u uf) find(x int) int {
	if u[x] != x {
		u[x] = u.find(u[x])
	}
	return u[x]
}

func (u uf) merge(from, to int) {
	u[u.find(from)] = u.find(to)
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 10
	var n, q, v, w, wt int
	Fscan(in, &n, &q)
	uf := [mx]uf{}
	for i := 1; i < mx; i++ {
		uf[i] = newUnionFind(n + 1)
	}

	ans := mx * (n - 1)
	for i := 0; i < n-1+q; i++ {
		Fscan(in, &v, &w, &wt)
		for ; wt < mx; wt++ {
			if !uf[wt].same(v, w) {
				ans--
				uf[wt].merge(v, w)
			}
		}
		if i >= n-1 {
			Fprintln(out, ans)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
