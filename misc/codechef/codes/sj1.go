package main

import (
	"bufio"
	. "fmt"
	"io"
)

func sj1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	for T := r(); T > 0; T-- {
		n := r()
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			v, w := r()-1, r()-1
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		a := make([]int, n)
		for i := range a {
			a[i] = r()
		}
		m := make([]int, n)
		for i := range m {
			m[i] = r()
		}

		g[0] = append(g[0], -1)
		allAns := make([]int, n)
		var dfs func(int, int, int)
		dfs = func(v, fa, s int) {
			s = gcd(s, a[v])
			if len(g[v]) == 1 {
				allAns[v] = m[v] - gcd(m[v], s)
			} else {
				allAns[v] = -1
			}
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, s)
				}
			}
		}
		dfs(0, -1, 0)
		for _, v := range allAns {
			if v >= 0 {
				Fprint(out, v, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { sj1(os.Stdin, os.Stdout) }
