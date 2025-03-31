package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
const eps4196 = 1e-8

type vec4196 struct{ x, y float64 }
type line4196 struct{ p1, p2 vec4196 }

func (a vec4196) add(b vec4196) vec4196 { return vec4196{a.x + b.x, a.y + b.y} }
func (a vec4196) sub(b vec4196) vec4196 { return vec4196{a.x - b.x, a.y - b.y} }
func (a vec4196) det(b vec4196) float64 { return a.x*b.y - a.y*b.x }
func (a vec4196) mul(k float64) vec4196 { return vec4196{a.x * k, a.y * k} }
func (a vec4196) polarAngle() float64      { return math.Atan2(a.y, a.x) }
func (a vec4196) onLeft(l line4196) bool   { return l.vec().det(a.sub(l.p1)) > eps4196 }
func (a line4196) vec() vec4196            { return a.p2.sub(a.p1) }
func (a line4196) point(t float64) vec4196 { return a.p1.add(a.vec().mul(t)) }
func (a line4196) intersection(b line4196) vec4196 {
	va, vb, u := a.vec(), b.vec(), a.p1.sub(b.p1)
	t := vb.det(u) / va.det(vb)
	return a.point(t)
}

func p4196(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type lp struct {
		l line4196
		p vec4196
	}
	halfPlanesIntersection := func(ls []line4196) float64 {
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
			if math.Abs(l.vec().det(q[len(q)-1].l.vec())) < eps4196 {
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

	ls := []line4196{}
	var n, m int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &m)
		ps := make([]vec4196, m)
		for i := range ps {
			Fscan(in, &ps[i].x, &ps[i].y)
		}
		for i := 1; i < m; i++ {
			ls = append(ls, line4196{ps[i-1], ps[i]})
		}
		ls = append(ls, line4196{ps[m-1], ps[0]})
	}
	Fprintf(out, "%.3f", halfPlanesIntersection(ls))
}

//func main() { p4196(os.Stdin, os.Stdout) }
