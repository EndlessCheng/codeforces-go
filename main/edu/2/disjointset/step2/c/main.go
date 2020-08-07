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
	Fscan(in, &n)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	for ; q > 0; q-- {
		Fscan(in, &op, &x, &y)
		if op == 1 {
			if x > y {
				x, y = y, x
			}
			fa[f(x)] = f(y)
		} else if op == 2 {
			fy := f(y)
			for {
				x = f(x)
				if x < fy {
					fa[x] = fy
					x++
				}
			}
		} else if f(x) == f(y) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
