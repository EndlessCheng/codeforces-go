package main

// https://space.bilibili.com/206214
func findClosest(x, y, z int) int {
	a := abs(x - z)
	b := abs(y - z)
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return 2
}

func abs(x int) int { if x < 0 { return -x }; return x }
