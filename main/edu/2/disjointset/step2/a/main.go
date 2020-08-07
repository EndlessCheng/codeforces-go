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

	var n, q, x int
	var s string
	Fscan(in, &n, &q)
	fa := make([]int, n+2)
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
		Fscan(in, &s, &x)
		if s[0] == '-' {
			fa[x] = x + 1
		} else {
			fx := f(x)
			if fx == n+1 {
				fx = -1
			}
			Fprintln(out, fx)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
