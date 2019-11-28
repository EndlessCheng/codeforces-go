package main

func minTimeToVisitAllPoints(points [][]int) int {
	n := len(points)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	for i := 1; i < n; i++ {
		pi := points[i]
		p0 := points[i-1]
		ans += max(abs(pi[0]-p0[0]), abs(pi[1]-p0[1]))
	}
	return ans
}
