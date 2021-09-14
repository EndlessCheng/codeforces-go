package main

// github.com/EndlessCheng/codeforces-go
func maximumUniqueSubarray(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	l := 0
	has := map[int]bool{}
	for i, v := range a {
		for has[v] {
			delete(has, a[l])
			l++
		}
		ans = max(ans, sum[i+1]-sum[l])
		has[v] = true
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
