package main

// https://space.bilibili.com/206214
func isReachableAtTime(sx, sy, fx, fy, t int) bool {
	if sx == fx && sy == fy {
		return t != 1
	}
	return max(abs(sx-fx), abs(sy-fy)) <= t
}

func abs(x int) int { if x < 0 { return -x }; return x }
func max(a, b int) int { if b > a { return b }; return a }
