package main

// github.com/EndlessCheng/codeforces-go
func minOperationsMaxProfit(a []int, boardingCost int, runningCost int) (ans int) {
	var left, maxS, s int
	do := func(i int) {
		c := 4
		if left < 4 {
			c = left
		}
		s += boardingCost*c - runningCost
		left -= c
		if s > maxS {
			maxS = s
			ans = i + 1
		}
	}
	for i, v := range a {
		left += v
		do(i)
	}
	for i := len(a); left > 0; i++ {
		do(i)
	}
	if maxS == 0 {
		return -1
	}
	return
}
