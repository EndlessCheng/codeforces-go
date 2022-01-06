package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF893C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, min int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n)
	var f func(int)
	f = func(v int) {
		vis[v] = true
		if a[v] < min {
			min = a[v]
		}
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	ans := int64(0)
	for i, b := range vis {
		if !b {
			min = 1e9
			f(i)
			ans += int64(min)
		}
	}
	Fprint(out, ans)
}

//func main() { CF893C(os.Stdin, os.Stdout) }
