package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1470D(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
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

o:
	for T := r(); T > 0; T-- {
		n := r()
		g := make([][]int, n)
		for m := r(); m > 0; m-- {
			v, w := r()-1, r()-1
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		color := make([]int8, n)
		var f func(int)
		f = func(v int) {
			if color[v] == 1 {
				todo := []int{}
				for _, w := range g[v] {
					if color[w] == 0 {
						color[w] = 2
						todo = append(todo, w)
					}
				}
				for _, w := range todo {
					f(w)
				}
			} else {
				for _, w := range g[v] {
					if color[w] == 0 {
						color[w] = 1
						f(w)
					}
				}
			}
		}
		color[0] = 1
		f(0)
		ans := []interface{}{}
		for i, c := range color {
			if c == 0 {
				Fprintln(out, "NO")
				continue o
			}
			if c == 1 {
				ans = append(ans, i+1)
			}
		}
		Fprintln(out, "YES")
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1470D(os.Stdin, os.Stdout) }
