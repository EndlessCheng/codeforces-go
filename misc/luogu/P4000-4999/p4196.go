package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
const eps196 = 1e-8

type vec196 struct{ x, y float64 }
type line196 struct{ p1, p2 vec196 }

func (a vec196) add(b vec196) vec196      { return vec196{a.x + b.x, a.y + b.y} }
func (a vec196) sub(b vec196) vec196      { return vec196{a.x - b.x, a.y - b.y} }
func (a vec196) det(b vec196) float64     { return a.x*b.y - a.y*b.x }
func (a vec196) mul(k float64) vec196     { return vec196{a.x * k, a.y * k} }
func (a vec196) polarAngle() float64     { return math.Atan2(a.y, a.x) }
func (a vec196) onLeft(l line196) bool   { return l.vec().det(a.sub(l.p1)) > eps196 }
func (a line196) vec() vec196            { return a.p2.sub(a.p1) }
func (a line196) point(t float64) vec196 { return a.p1.add(a.vec().mul(t)) }
func (a line196) intersection(b line196) vec196 {
	va, vb, u := a.vec(), b.vec(), a.p1.sub(b.p1)
	t := vb.det(u) / va.det(vb)
	return a.point(t)
}

func p4196(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type lp struct {
		l line196
		p vec196
	}
	halfPlanesIntersection := func(ls []line196) float64 {
		sort.Slice(ls, func(i, j int) bool { return ls[i].vec().polarAngle() < ls[j].vec().polarAngle() })
		q := []lp{{l: ls[0]}}
		for i := 1; i < len(ls); i++ {
			l := ls[i]
			for len(q) > 1 && !q[len(q)-2].p.onLeft(l) {
				q = q[:len(q)-1]
			}
			for len(q) > 1 && !q[0].p.onLeft(l) {
				q = q[1:]
			}
			if math.Abs(l.vec().det(q[len(q)-1].l.vec())) < eps196 {
				if l.p1.onLeft(q[len(q)-1].l) {
					q[len(q)-1].l = l
				}
			} else {
				q = append(q, lp{l: l})
			}
			if len(q) > 1 {
				q[len(q)-2].p = q[len(q)-2].l.intersection(q[len(q)-1].l)
			}
		}
		for len(q) > 1 && !q[len(q)-2].p.onLeft(q[0].l) {
			q = q[:len(q)-1]
		}
		q[len(q)-1].p = q[len(q)-1].l.intersection(q[0].l)
		area := 0.
		p0 := q[0].p
		for i := 2; i < len(q); i++ {
			area += q[i-1].p.sub(p0).det(q[i].p.sub(p0))
		}
		return area / 2
	}

	ls := []line196{}
	var n, m int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &m)
		ps := make([]vec196, m)
		for i := range ps {
			Fscan(in, &ps[i].x, &ps[i].y)
		}
		for i := 1; i < m; i++ {
			ls = append(ls, line196{ps[i-1], ps[i]})
		}
		ls = append(ls, line196{ps[m-1], ps[0]})
	}
	Fprintf(out, "%.3f", halfPlanesIntersection(ls))
}

//func main() { p4196(os.Stdin, os.Stdout) }
