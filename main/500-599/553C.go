package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF553C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7

	var n, m, v, w, t int
	Fscan(in, &n, &m)
	g := [2][][]int{make([][]int, n), make([][]int, n)}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &t)
		v--
		w--
		g[t][v] = append(g[t][v], w)
		g[t][w] = append(g[t][w], v)
	}

	colors := make([]int8, n)
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[1][v] {
			if colors[w] == -c || colors[w] == 0 && !f(w, c) {
				return false
			}
		}
		for _, w := range g[0][v] {
			if colors[w] == c || colors[w] == 0 && !f(w, -c) {
				return false
			}
		}
		return true
	}
	ans := (mod + 1) / 2
	for i, c := range colors {
		if c == 0 {
			if !f(i, 1) {
				Fprint(out, 0)
				return
			}
			ans = ans * 2 % mod
		}
	}
	Fprint(out, ans)
}

//func main() { CF553C(os.Stdin, os.Stdout) }
