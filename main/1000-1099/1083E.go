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
	var n, ans int
	Fscan(in, &n)
	type tuple struct{ x, y, a int }
	a := make([]tuple, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].a)
	}
	slices.SortFunc(a, func(a, b tuple) int { return a.x - b.x })

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
