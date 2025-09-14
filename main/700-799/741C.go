package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf741C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	x := make([]int, n+1)
	y := make([]int, n+1)
	g := make([][]int, n*2+1)
	add := func(v, w int) {
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &x[i], &y[i])
		add(x[i], y[i])
		add(i*2-1, i*2)
	}

	cs := make([]int8, len(g))
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		cs[v] = c
		for _, w := range g[v] {
			if cs[w] == c || cs[w] == 0 && !f(w, 3^c) {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n*2; i++ {
		if cs[i] == 0 && !f(i, 1) {
			Fprint(out, -1)
			return
		}
	}

	for i := 1; i <= n; i++ {
		Fprintln(out, cs[x[i]], cs[y[i]])
	}
}

//func main() { cf741C(bufio.NewReader(os.Stdin), os.Stdout) }
