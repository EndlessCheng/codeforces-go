package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF219D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	type edge struct{ to, ord int }
	g := make([][]edge, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], edge{w, 1})
		g[w] = append(g[w], edge{v, -1})
	}
	inv := make([]int, n)
	var f func(v, fa int)
	f = func(v, fa int) {
		for _, e := range g[v] {
			if w := e.to; w != fa {
				if e.ord == -1 {
					inv[0]++
				}
				f(w, v)
			}
		}
	}
	f(0, -1)
	f = func(v, fa int) {
		for _, e := range g[v] {
			if w := e.to; w != fa {
				inv[w] = inv[v] + e.ord
				f(w, v)
			}
		}
	}
	f(0, -1)
	min := n
	for _, v := range inv {
		if v < min {
			min = v
		}
	}
	Fprintln(out, min)
	for i, v := range inv {
		if v == min {
			Fprint(out, i+1, " ")
		}
	}
}

func main() { CF219D(os.Stdin, os.Stdout) }
