package main

import "math"

// github.com/EndlessCheng/codeforces-go
func checkOverlap(r, ox, oy, x1, y1, x2, y2 int) (ans bool) {
	rx, ry := float64(x1+x2)/2, float64(y1+y2)/2
	hx, hy := float64(x2-x1)/2, float64(y2-y1)/2
	x, y := math.Abs(float64(ox)-rx), math.Abs(float64(oy)-ry)
	x, y = math.Max(x-hx, 0), math.Max(y-hy, 0)
	return x*x+y*y < float64(r*r)+1e-8
}
