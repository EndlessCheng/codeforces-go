package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1006E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, k int
	Fscan(in, &n, &q)
	g := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		Fscan(in, &v)
		g[v] = append(g[v], i)
	}
	pos := make([][2]int, n+1)
	t := make([]int, 0, n)
	var f func(int)
	f = func(v int) {
		pos[v][0] = len(t)
		t = append(t, v)
		for _, w := range g[v] {
			f(w)
		}
		pos[v][1] = len(t)
	}
	f(1)
	for ; q > 0; q-- {
		Fscan(in, &v, &k)
		if p := pos[v][0] + k - 1; p < pos[v][1] {
			Fprintln(out, t[p])
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1006E(os.Stdin, os.Stdout) }
