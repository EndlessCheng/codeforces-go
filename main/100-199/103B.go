package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF103B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, c int
	Fscan(in, &n, &m)
	if m != n {
		Fprint(out, "NO")
		return
	}
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make([]bool, len(g))
	var f func(int)
	f = func(v int) {
		vis[v] = true
		c++
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	f(1)
	if c == n {
		Fprint(out, "FHTAGN!")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF103B(os.Stdin, os.Stdout) }
