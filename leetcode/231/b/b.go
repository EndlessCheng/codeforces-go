package main

// github.com/EndlessCheng/codeforces-go
func minElements(a []int, limit, goal int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return (abs(sum-goal) + limit - 1) / limit
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
