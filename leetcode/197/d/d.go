package main

import (
	"math"
	"math/rand"
)

// github.com/EndlessCheng/codeforces-go
func getMinDistSum(positions [][]int) (ans float64) {
	n := len(positions)
	ps, xx, yy := make([][2]float64, n), 0, 0
	for i, p := range positions {
		ps[i] = [2]float64{float64(p[0]), float64(p[1])}
		xx += p[0]
		yy += p[1]
	}

	sumD := func(x, y float64) (l float64) {
		for _, p := range ps {
			l += math.Hypot(p[0]-x, p[1]-y)
		}
		return l
	}
	x := float64(xx) / float64(n)
	y := float64(yy) / float64(n)
	ans = sumD(x, y)
	for t := 1e2; t > 1e-8; t *= 0.99 {
		dx, dy := 0.0, 0.0
		for _, p := range ps {
			d := math.Hypot(x-p[0], y-p[1])
			dx += (p[0] - x) / d
			dy += (p[1] - y) / d
		}
		xx, yy := x+t*dx, y+t*dy
		if s := sumD(xx, yy); s < ans || math.Exp((ans-s)/t) > rand.Float64() {
			ans, x, y = s, xx, yy
		}
	}
	return
}
