package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2044H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 4096)
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

	for range r() {
		n, q := r(), r()
		s := make([][][3]int, n+1)
		for i := range s {
			s[i] = make([][3]int, n+1)
		}
		for i := range n {
			for j := range n {
				v := r()
				s[i+1][j+1][0] = s[i+1][j][0] + s[i][j+1][0] - s[i][j][0] + v
				s[i+1][j+1][1] = s[i+1][j][1] + s[i][j+1][1] - s[i][j][1] + v*i
				s[i+1][j+1][2] = s[i+1][j][2] + s[i][j+1][2] - s[i][j][2] + v*(j+1)
			}
		}
		for range q {
			r1, c1, r2, c2 := r()-1, r()-1, r(), r()
			query := func(t int) int { return s[r2][c2][t] - s[r2][c1][t] - s[r1][c2][t] + s[r1][c1][t] }
			s0, s1, s2 := query(0), query(1), query(2)
			Fprint(out, (c2-c1)*(s1-r1*s0)+s2-c1*s0, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2044H(os.Stdin, os.Stdout) }
