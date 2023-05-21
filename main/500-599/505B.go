package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF505B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, c, q int
	Fscan(in, &n, &m)
	fa := make([]map[int]int, m)
	for i := range fa {
		fa[i] = map[int]int{}
	}
	var find func(int, int) int
	find = func(c, x int) int {
		f, ok := fa[c][x]
		if !ok {
			fa[c][x] = x
			f = x
		}
		if f != x {
			fa[c][x] = find(c, f)
			return fa[c][x]
		}
		return x
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &c)
		c--
		v = find(c, v)
		w = find(c, w)
		fa[c][v] = w
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &w)
		cnt := 0
		for i := range fa {
			if find(i, v) == find(i, w) {
				cnt++
			}
		}
		Fprintln(out, cnt)
	}
}

//func main() { CF505B(os.Stdin, os.Stdout) }
