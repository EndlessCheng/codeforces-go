package main

// github.com/EndlessCheng/codeforces-go
func nearestValidPoint(x, y int, points [][]int) int {
	ans := -1
	mi := int(1e9)
	for i, p := range points {
		if p[0] == x || p[1] == y {
			if d := abs(p[0]-x) + abs(p[1]-y); d < mi {
				mi, ans = d, i
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
