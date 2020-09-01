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
	fa := make([]int, n+1)
	next := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		next[i] = i + 1
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[f(from)] = f(to) }
	for ; q > 0; q-- {
		Fscan(in, &op, &x, &y)
		if op == 1 {
			merge(x, y)
		} else if op == 2 {
			for i := x + 1; i <= y; {
				merge(i-1, i)
				i, next[i] = next[i], next[y]
			}
		} else if f(x) == f(y) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
