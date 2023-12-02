package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF369C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, t int
	Fscan(in, &n)
	type nb struct{ to, t int }
	g := make([][]nb, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &t)
		g[v] = append(g[v], nb{w, t})
		g[w] = append(g[w], nb{v, t})
	}

	ans := []int{}
	var f func(int, int) bool
	f = func(v, fa int) (choose bool) {
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			c := f(w, v)
			if !c && e.t == 2 {
				ans = append(ans, w)
				c = true
			}
			choose = choose || c
		}
		return
	}
	f(1, 0)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF369C(os.Stdin, os.Stdout) }
