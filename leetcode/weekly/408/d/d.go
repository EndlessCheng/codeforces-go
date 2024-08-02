package main

// https://space.bilibili.com/206214
// 判断点 (x,y) 是否在圆 (ox,oy,r) 内
func inCircle(ox, oy, r, x, y int) bool {
	return (ox-x)*(ox-x)+(oy-y)*(oy-y) <= r*r
}

func canReachCorner(X, Y int, circles [][]int) bool {
	vis := make([]bool, len(circles))
	var dfs func(int) bool
	dfs = func(i int) bool {
		x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
		// 圆 i 是否与矩形右边界/下边界相交相切
		if y1 <= Y && abs(x1-X) <= r1 || x1 <= X && y1 <= r1 || x1 > X && inCircle(x1, y1, r1, X, 0) {
			return true
		}
		vis[i] = true
		for j, c := range circles {
			x2, y2, r2 := c[0], c[1], c[2]
			// 在两圆相交相切的前提下，点 A 是否严格在矩形内
			if !vis[j] && (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= (r1+r2)*(r1+r2) &&
				x1*r2+x2*r1 < (r1+r2)*X &&
				y1*r2+y2*r1 < (r1+r2)*Y &&
				dfs(j) {
				return true
			}
		}
		return false
	}
	for i, c := range circles {
		x, y, r := c[0], c[1], c[2]
		if inCircle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
			inCircle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
			// 圆 i 是否与矩形上边界/左边界相交相切
			!vis[i] && (x <= X && abs(y-Y) <= r || y <= Y && x <= r || y > Y && inCircle(x, y, r, 0, Y)) && dfs(i) {
			return false
		}
	}
	return true
}

func abs(x int) int { if x < 0 { return -x }; return x }
