package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF982C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, ans int
	Fscan(in, &n)
	if n&1 == 1 {
		Fprint(_w, -1)
		return
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var f func(int, int) int
	f = func(v, fa int) int {
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v)
			}
		}
		if fa != -1 && sz&1 == 0 {
			ans++
		}
		return sz
	}
	f(0, -1)
	Fprint(_w, ans)
}

//func main() { CF982C(os.Stdin, os.Stdout) }
