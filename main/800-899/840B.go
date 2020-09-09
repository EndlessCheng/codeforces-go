package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF840B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	neg, s := -1, int8(0)
	d := make([]int8, n)
	for i := range d {
		Fscan(in, &d[i])
		if d[i] < 0 {
			d[i] = 0
			neg = i
		} else {
			s ^= d[i]
		}
	}
	if neg < 0 {
		if s > 0 {
			Fprint(out, -1)
			return
		}
	} else {
		d[neg] = s
	}

	type nb struct{ to, i int }
	g := make([][]nb, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], nb{w, i})
		g[w] = append(g[w], nb{v, i})
	}
	ans := []interface{}{}
	vis := make([]bool, n)
	var f func(int) int8
	f = func(v int) int8 {
		vis[v] = true
		s := d[v]
		for _, e := range g[v] {
			if w := e.to; !vis[w] && f(w) > 0 {
				ans = append(ans, e.i)
				s ^= 1
			}
		}
		return s
	}
	f(0)
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { CF840B(os.Stdin, os.Stdout) }
