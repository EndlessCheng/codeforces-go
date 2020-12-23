package main

import (
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type vec struct{ x, y float64 }

func (a vec) sub(b vec) vec      { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) float64  { return a.x*b.y - a.y*b.x }
func (a vec) len2() float64      { return a.x*a.x + a.y*a.y }
func (a vec) dis2(b vec) float64 { return a.sub(b).len2() }
func (a vec) less(b vec) bool    { return a.x < b.x || a.x == b.x && a.y < b.y }

func solve(n int, a [][]int) float64 {
	convexHull := func(ps []vec) []vec {
		n := len(ps)
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
		ch := make([]vec, 0, 2*n)
		for _, p := range ps {
			for {
				sz := len(ch)
				if sz <= 1 || ch[sz-1].sub(ch[sz-2]).det(p.sub(ch[sz-1])) > 0 {
					break
				}
				ch = ch[:sz-1]
			}
			ch = append(ch, p)
		}
		downSize := len(ch)
		for i := n - 2; i >= 0; i-- {
			p := ps[i]
			for {
				sz := len(ch)
				if sz <= downSize || ch[sz-1].sub(ch[sz-2]).det(p.sub(ch[sz-1])) > 0 {
					break
				}
				ch = ch[:sz-1]
			}
			ch = append(ch, p)
		}
		return ch[:len(ch)-1]
	}
	rotatingCalipers := func(ps []vec) float64 {
		qs := convexHull(ps)
		n := len(qs)
		if n == 2 {
			return qs[0].dis2(qs[1])
		}
		var p1, p2 vec
		i, j := 0, 0
		for k, p := range qs {
			if !qs[i].less(p) {
				i = k
			}
			if qs[j].less(p) {
				j = k
			}
		}
		maxDis2 := 0.
		i0, j0 := i, j
		for i != j0 || j != i0 {
			if d2 := qs[i].sub(qs[j]).len2(); d2 > maxDis2 {
				maxDis2 = d2
				p1, p2 = qs[i], qs[j]
			}
			if qs[(i+1)%n].sub(qs[i]).det(qs[(j+1)%n].sub(qs[j])) < 0 {
				i = (i + 1) % n
			} else {
				j = (j + 1) % n
			}
		}
		return p1.dis2(p2)
	}
	dis := func(t float64) float64 {
		b := make([]vec, n)
		for i, p := range a {
			b[i].x = float64(p[0]) + t*float64(p[2])
			b[i].y = float64(p[1]) + t*float64(p[3])
		}
		return rotatingCalipers(b)
	}
	ternarySearch := func(l, r float64, f func(x float64) float64) float64 {
		step := int(math.Log((r-l)/1e-8) / math.Log(1.5))
		for ; step > 0; step-- {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2
			} else {
				l = m1
			}
		}
		return (l + r) / 2
	}
	return math.Sqrt(dis(ternarySearch(0, 1e5, dis)))
}
