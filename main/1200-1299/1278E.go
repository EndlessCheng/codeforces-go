package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1278E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, lastR int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	ans := make([][2]int, n+1)
	var f func(v, fa, r int)
	f = func(v, fa, r int) {
		lastR = r
		c := 0
		for _, w := range g[v] {
			if w != fa {
				c++
				ans[w] = [2]int{r - c, lastR + len(g[w])}
				f(w, v, ans[w][1])
			}
		}
	}
	f(1, 0, 2+len(g[1]))
	Fprintln(out, 1, 2+len(g[1]))
	for _, s := range ans[2:] {
		Fprintln(out, s[0], s[1])
	}
}

//func main() { CF1278E(os.Stdin, os.Stdout) }
