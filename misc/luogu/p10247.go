package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func p10247(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, s, t int
	Fscan(in, &n, &m)
	es := make([]struct{ x, y int }, m)
	deg := make([]int, n+1)
	for i := range es {
		Fscan(in, &es[i].x, &es[i].y)
		deg[es[i].x]++
		deg[es[i].y]++
	}

	// 特判菊花图
	for _, d := range deg {
		if d == m {
			Fprintln(out, strings.Repeat("0 ", m))
			return
		}
	}

	ok := func(i, j int) bool {
		a, b := es[i].x, es[i].y
		c, d := es[j].x, es[j].y
		return a != c && a != d && b != c && b != d
	}
	// 寻找两条不相邻的边
	for i := range es {
		for j := i + 1; j < m; j++ {
			if ok(i, j) {
				s, t = i, j
				goto next
			}
		}
	}
	Fprintln(out, strings.Repeat("0 ", m))
	return

next:
	ans := make([]int, m)
	for i := range ans {
		if ok(i, s) {
			ans[i] = s + 1
		} else if ok(i, t) {
			ans[i] = t + 1
		} else {
			for j := range es {
				if ok(i, j) {
					ans[i] = j + 1
				}
			}
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
	Fprintln(out)
}

//func main() { p10247(os.Stdin, os.Stdout) }
