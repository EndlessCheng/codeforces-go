package main

// github.com/EndlessCheng/codeforces-go
func numPairsDivisibleBy60(time []int) int {
	c := [60]int{}
	for _, v := range time {
		c[v%60]++
	}
	ans := c[0]*(c[0]-1)/2 + c[30]*(c[30]-1)/2
	for i := 1; i < 30; i++ {
		ans += c[i] * c[60-i]
	}
	return ans
}
