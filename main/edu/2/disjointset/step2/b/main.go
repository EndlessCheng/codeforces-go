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

	var n, x int
	Fscan(in, &n)
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
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		fx := f(x)
		if fx == n+1 {
			fx = f(1)
		}
		Fprint(out, fx, " ")
		fa[fx] = fx + 1
	}
}

func main() { run(os.Stdin, os.Stdout) }
