package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type vec3628 struct{ x, y int }

func (a vec3628) sub(b vec3628) vec3628 { return vec3628{a.x - b.x, a.y - b.y} }
func (a vec3628) dot(b vec3628) int     { return a.x*b.x + a.y*b.y }
func (a vec3628) det(b vec3628) int     { return a.x*b.y - a.y*b.x }

func p3628(in io.Reader, out io.Writer) {
	buf := make([]byte, 4096)
	_i, _n := 0, 0
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
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	rn := func() (x int) {
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	n, a, b, c := rd(), rn(), rn(), rn()
	q := []vec3628{{}}
	for s := 0; ; n-- {
		s += rd()
		pi := vec3628{-2 * a * s, 1}
		for len(q) > 1 && pi.dot(q[0]) <= pi.dot(q[1]) {
			q = q[1:]
		}
		f := pi.dot(q[0]) + a*s*s + b*s + c
		if n == 1 {
			Fprint(out, f)
			return
		}
		vj := vec3628{s, a*s*s - b*s + f}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(vj.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, vj)
	}
}

//func main() { p3628(os.Stdin, os.Stdout) }
