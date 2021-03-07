package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1131F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)
	fa := make([]int, n+1)
	vs := make([][]interface{}, n+1)
	for i := range fa {
		fa[i] = i
		vs[i] = []interface{}{i}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for i := 1; i < n; i++ {
		Fscan(in, &x, &y)
		x, y = find(x), find(y)
		if len(vs[x]) < len(vs[y]) {
			x, y = y, x
		}
		vs[x] = append(vs[x], vs[y]...)
		fa[y] = x
	}
	Fprintln(out, vs[find(1)]...)
}

//func main() { CF1131F(os.Stdin, os.Stdout) }
