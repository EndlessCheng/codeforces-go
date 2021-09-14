package main

// github.com/EndlessCheng/codeforces-go
func getMinDistance(a []int, tar, start int) int {
	ans := int(1e9)
	for i, v := range a {
		if v == tar {
			ans = min(ans, abs(i-start))
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
