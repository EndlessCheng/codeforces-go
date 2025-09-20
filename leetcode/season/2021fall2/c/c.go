package main

import (
	"math"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
type vec struct{ x, y int }

func (a *vec) adds(b vec) { a.x += b.x; a.y += b.y }
func (a *vec) subs(b vec) { a.x -= b.x; a.y -= b.y }

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	a := make([]vec, n)
	a[0].y = 1
	for i, v := range finalCnt {
		a[i+1].x = v
	}

	for _, p := range slices.Backward(plans) {
		i := p[1]
		if p[0] == 1 {
			a[i].x *= 2
			a[i].y *= 2
		} else if p[0] == 2 {
			for _, w := range g[i] {
				a[w].subs(a[i])
			}
		} else {
			for _, w := range g[i] {
				a[w].adds(a[i])
			}
		}
	}

	s := vec{}
	for _, v := range a {
		s.adds(v)
	}
	x := (int(totalNum) - s.x) / s.y

	ans := make([]int, n)
	for i, p := range a {
		ans[i] = p.x + p.y*x
	}
	return ans
}

// 虚数做法
func volunteerDeployment1(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	a := make([]complex128, n)
	a[0] = 1i
	for i, v := range finalCnt {
		a[i+1] = complex(float64(v), 0)
	}

	for _, p := range slices.Backward(plans) {
		i := p[1]
		if p[0] == 1 {
			a[i] *= 2
		} else if p[0] == 2 {
			for _, w := range g[i] {
				a[w] -= a[i]
			}
		} else {
			for _, w := range g[i] {
				a[w] += a[i]
			}
		}
	}

	s := 0i
	for _, v := range a {
		s += v
	}
	x := int(math.Round((float64(totalNum) - real(s)) / imag(s)))

	ans := make([]int, n)
	for i, c := range a {
		ans[i] = int(real(c)) + int(imag(c))*x
	}
	return ans
}