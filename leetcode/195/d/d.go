package main

func findMaxValueOfEquation(points [][]int, k int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	f := func(p []int) int { return p[1] - p[0] }

	ans = -1e18
	idQ := make([]int, len(points))
	l, r := 0, 0
	for i, p := range points {
		for ; l < r && p[0]-points[idQ[l]][0] > k; l++ {
		}
		if l < r {
			ans = max(ans, p[0]+p[1]+f(points[idQ[l]]))
		}
		for ; l < r && f(points[idQ[r-1]]) <= f(p); r-- {
		}
		idQ[r] = i
		r++
	}
	return
}
