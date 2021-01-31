package main

// github.com/EndlessCheng/codeforces-go
func countBalls(lowLimit int, highLimit int) (ans int) {
	c := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		s := 0
		for x := i; x > 0; x /= 10 {
			s += x % 10
		}
		c[s]++
		ans = max(ans, c[s])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
