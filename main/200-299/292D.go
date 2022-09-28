package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF292D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, l, r int
	Fscan(in, &n, &m)
	fa := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	pre := make([][]int, m+1)
	for i := range fa {
		fa[i] = i
	}
	pre[0] = append([]int{}, fa...)
	e := make([]struct{ x, y int }, m)
	for i := range e {
		Fscan(in, &e[i].x, &e[i].y)
		e[i].x--
		e[i].y--
		fa[find(e[i].x)] = find(e[i].y)
		pre[i+1] = append([]int{}, fa...)
	}

	suf := make([][]int, m+1)
	for i := range fa {
		fa[i] = i
	}
	suf[m] = append([]int{}, fa...)
	for i := m - 1; i > 0; i-- {
		fa[find(e[i].x)] = find(e[i].y)
		suf[i] = append([]int{}, fa...)
	}

	for Fscan(in, &k); k > 0; k-- {
		Fscan(in, &l, &r)
		fa = append([]int{}, pre[l-1]...)
		for i, f := range suf[r] {
			fa[find(fa[i])] = find(f)
		}
		c := 0
		for i, f := range fa {
			if i == f {
				c++
			}
		}
		Fprintln(out, c)
	}
}

//func main() { CF292D(os.Stdin, os.Stdout) }
