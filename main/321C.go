package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF321C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	type node struct{ dep, fa int }
	nodes := make([]node, n)
	var _f func(v, fa, d int)
	_f = func(v, fa, d int) {
		nodes[v] = node{d, fa}
		for _, w := range g[v] {
			if w != fa {
				_f(w, v, d+1)
			}
		}
	}
	_f(0, -1, 0)

	used := make([]bool, n)
	size := make([]int, n)
	var calcSize func(v, fa int) int
	calcSize = func(v, fa int) int {
		sz := 1
		for _, w := range g[v] {
			if w != fa && !used[w] {
				sz += calcSize(w, v)
			}
		}
		size[v] = sz
		return sz
	}
	var compSize int
	var find func(v, fa int) int
	find = func(v, fa int) int {
		for _, w := range g[v] {
			if w != fa && !used[w] && size[w] > compSize>>1 {
				return find(w, v)
			}
		}
		return v
	}

	ans := make([]byte, n)
	var f func(int, byte)
	f = func(v int, rk byte) {
		calcSize(v, -1)
		compSize = size[v]
		ct := find(v, -1)
		used[ct] = true
		ans[ct] = rk
		for _, w := range g[ct] {
			if !used[w] {
				f(w, rk+1)
			}
		}
	}
	f(0, 'A')
	for _, b := range ans {
		Fprint(out, string(b)+" ")
	}
}

//func main() { CF321C(os.Stdin, os.Stdout) }
