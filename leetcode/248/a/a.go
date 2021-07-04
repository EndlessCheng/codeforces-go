package main

// github.com/EndlessCheng/codeforces-go
func buildArray(a []int) []int {
	ans := make([]int, len(a))
	for i, v := range a {
		ans[i] = a[v]
	}
	return ans
}
