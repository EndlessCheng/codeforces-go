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
		if Fscan(in, &s, &x, &y); s[0] == 'u' {
			fa[f(x)] = f(y)
		} else if f(x) == f(y) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
