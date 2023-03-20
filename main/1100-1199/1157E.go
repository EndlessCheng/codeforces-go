package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1157E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}

	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	mergeNext := func(x int) { fa[find(x)] = find(x + 1) }

	for i, c := range cnt {
		if c == 0 {
			mergeNext(i)
		}
	}
	for _, x := range a {
		y := find((n - x) % n)
		if y == n {
			y = find(0)
		}
		cnt[y]--
		if cnt[y] == 0 {
			mergeNext(y)
		}
		Fprint(out, (x+y)%n, " ")
	}
}

//func main() { CF1157E(os.Stdin, os.Stdout) }
