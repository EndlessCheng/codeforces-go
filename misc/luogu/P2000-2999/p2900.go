package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type vec900 struct{ x, y int }

func (a vec900) sub(b vec900) vec900 { return vec900{a.x - b.x, a.y - b.y} }
func (a vec900) dot(b vec900) int    { return a.x*b.x + a.y*b.y }
func (a vec900) det(b vec900) int    { return a.x*b.y - a.y*b.x }

func p2900(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ x, y int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Or(a.x-b.x, a.y-b.y) })

	st := a[:0]
	for _, p := range a {
		for len(st) > 0 && st[len(st)-1].y <= p.y {
			st = st[:len(st)-1]
		}
		st = append(st, p)
	}

	q := []vec900{}
	f := 0
	for _, p := range st {
		vj := vec900{-p.y, f}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(vj.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, vj)

		pi := vec900{-p.x, 1}
		for len(q) > 1 && pi.dot(q[0]) >= pi.dot(q[1]) {
			q = q[1:]
		}
		f = pi.dot(q[0])
	}
	Fprint(out, f)
}

//func main() { p2900(bufio.NewReader(os.Stdin), os.Stdout) }
