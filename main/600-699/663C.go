package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF663C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	var s string
	Fscan(in, &n, &m)
	type nb struct{ to, c int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &s)
		v--
		w--
		c := 1
		if s == "R" {
			c = 2
		}
		g[v] = append(g[v], nb{w, c})
		g[w] = append(g[w], nb{v, c})
	}
	ans := [3][]int{}
	for choose := 1; choose < 3; choose++ {
		vs := [3][]int{}
		color := make([]int, n)
		var f func(int, int) bool
		f = func(v, c int) bool {
			vs[c] = append(vs[c], v)
			color[v] = c
			for _, e := range g[v] {
				if w := e.to; color[w] == c^e.c^choose^3 || color[w] == 0 && !f(w, c^e.c^choose) {
					return false
				}
			}
			return true
		}
		for i, c := range color {
			if c > 0 {
				continue
			}
			vs = [3][]int{}
			if !f(i, 1) {
				ans[choose] = make([]int, n+1)
				break
			}
			if len(vs[1]) > len(vs[2]) {
				ans[choose] = append(ans[choose], vs[1]...)
			} else {
				ans[choose] = append(ans[choose], vs[2]...)
			}
		}
	}
	if len(ans[1]) > len(ans[2]) {
		ans[1] = ans[2]
	}
	if len(ans[1]) > n {
		Fprint(out, -1)
		return
	}
	Fprintln(out, len(ans[1]))
	for _, v := range ans[1] {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF663C(os.Stdin, os.Stdout) }
