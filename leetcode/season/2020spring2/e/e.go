package main

import (
	"sort"
)

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }

type pair struct {
	vec
	i int
}

func visitOrder(points [][]int, s string) (ans []int) {
	ans = make([]int, 0, len(points))
	ps := make([]pair, len(points))
	for i, p := range points {
		ps[i] = pair{vec{p[0], p[1]}, i}
	}

	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
	p0 := ps[0]
	ps = ps[1:]
	ans = append(ans, p0.i)

	leftMostVec := func(p0 vec) (idx int) {
		for i, p := range ps {
			if ps[idx].sub(p0).det(p.sub(p0)) > 0 {
				idx = i
			}
		}
		return
	}
	rightMostVec := func(p0 vec) (idx int) {
		for i, p := range ps {
			if ps[idx].sub(p0).det(p.sub(p0)) < 0 {
				idx = i
			}
		}
		return
	}
	for i, b := range s {
		if b == 'L' {
			i = rightMostVec(p0.vec)
		} else {
			i = leftMostVec(p0.vec)
		}
		p0, ps[i] = ps[i], ps[0]
		ps = ps[1:]
		ans = append(ans, p0.i)
	}
	ans = append(ans, ps[0].i)
	return
}
