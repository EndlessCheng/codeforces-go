package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF538E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, v, w, leaf int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
	}

	var f func(int, bool) int
	f = func(v int, tp bool) int {
		if g[v] == nil {
			leaf++
			return 1
		}
		if tp {
			res := 0
			for _, w := range g[v] {
				res += f(w, !tp)
			}
			return res
		} else {
			res := n + 1
			for _, w := range g[v] {
				res = min(res, f(w, !tp))
			}
			return res
		}
	}
	mx, mi := f(1, false), f(1, true)
	Fprint(out, leaf/2+1-mx, mi)
}

//func main() { CF538E(os.Stdin, os.Stdout) }
