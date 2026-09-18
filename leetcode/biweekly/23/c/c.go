package main

import "math"

// github.com/EndlessCheng/codeforces-go
func checkOverlap1(r, ox, oy, x1, y1, x2, y2 int) (ans bool) {
	rx, ry := float64(x1+x2)/2, float64(y1+y2)/2
	hx, hy := float64(x2-x1)/2, float64(y2-y1)/2
	x, y := math.Abs(float64(ox)-rx), math.Abs(float64(oy)-ry)
	x, y = math.Max(x-hx, 0), math.Max(y-hy, 0)
	return x*x+y*y < float64(r*r)+1e-8
}

func checkOverlap(radius, xCenter, yCenter, x1, y1, x2, y2 int) bool {
	// 找到在矩形中的到圆心 (xCenter, yCenter) 最近的点 (x, y)
	x := max(x1, min(xCenter, x2))
	y := max(y1, min(yCenter, y2))

	// 判断 (x, y) 是否在圆中
	return (x-xCenter)*(x-xCenter)+(y-yCenter)*(y-yCenter) <= radius*radius
}
