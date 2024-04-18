package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf371D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	cap := make([]int, n)
	for i := range cap {
		Fscan(in, &cap[i])
	}
	water := make([]int, n)

	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	var m, op, i, x int
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &op, &i)
		i--
		if op == 2 {
			Fprintln(out, water[i])
			continue
		}
		Fscan(in, &x)
		for i = find(i); i < n; i = find(i) {
			if water[i]+x < cap[i] {
				water[i] += x
				break
			}
			x -= cap[i] - water[i]
			water[i] = cap[i]
			fa[i] = i + 1
		}
	}
}

//func main() { cf371D(os.Stdin, os.Stdout) }
