package main

import (
	. "fmt"
	"io"
	"math/big"
	"slices"
)

// https://github.com/EndlessCheng
type vec83 struct{ x, y int }

func (a vec83) sub(b vec83) vec83 { return vec83{a.x - b.x, a.y - b.y} }
func (a vec83) dot(b vec83) int   { return a.x*b.x + a.y*b.y }
func (a vec83) det(b vec83) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func cf1083E(in io.Reader, out io.Writer) {
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

	type tuple struct{ x, y, a int }
	a := make([]tuple, rd())
	for i := range a {
		a[i] = tuple{rd(), rd(), rd()}
	}
	slices.SortFunc(a, func(a, b tuple) int { return a.x - b.x })

	ans := 0
	q := []vec83{{}}
	for _, t := range a {
		pi := vec83{-t.y, 1}
		for len(q) > 1 && pi.dot(q[0]) <= pi.dot(q[1]) {
			q = q[1:]
		}
		f := pi.dot(q[0]) + t.x*t.y - t.a
		ans = max(ans, f)

		vj := vec83{t.x, f}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(vj.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, vj)
	}
	Fprint(out, ans)
}

//func main() { cf1083E(bufio.NewReader(os.Stdin), os.Stdout) }
