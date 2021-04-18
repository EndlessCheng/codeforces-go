package main

// github.com/EndlessCheng/codeforces-go
func getMaximumXor(a []int, k int) []int {
	n := len(a)
	ans := make([]int, n)
	s, m := 0, 1<<k-1
	for i, v := range a {
		s ^= v
		ans[n-1-i] = s&m ^ m
	}
	return ans
}
