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

	var n, q, x, y int
	var s string
	Fscan(in, &n, &q)
	fa := make([]int, n+1)
	sz := make([]int, n+1)
	min := make([]int, n+1)
	max := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
		min[i] = i
		max[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	for ; q > 0; q-- {
		if Fscan(in, &s, &x); s[0] == 'u' {
			Fscan(in, &y)
			fx, fy := f(x), f(y)
			if fx == fy {
				continue
			}
			fa[fx] = fy
			sz[fy] += sz[fx]
			if min[fx] < min[fy] {
				min[fy] = min[fx]
			}
			if max[fx] > max[fy] {
				max[fy] = max[fx]
			}
		} else {
			fx := f(x)
			Fprintln(out, min[fx], max[fx], sz[fx])
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
