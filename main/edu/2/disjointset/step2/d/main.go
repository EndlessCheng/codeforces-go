package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, x, y int
	Fscan(in, &n, &q)
	dis := make([]int, n+1)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			ffx := f(fa[x])
			dis[x] += dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	for ; q > 0; q-- {
		Fscan(in, &op, &x)
		if op == 1 {
			Fscan(in, &y)
			if fx, fy := f(x), f(y); fx != fy {
				dis[fx] = 1 + dis[y] - dis[x]
				fa[fx] = fy
			}
		} else {
			f(x)
			Fprintln(out, dis[x])
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
