package main

// github.com/EndlessCheng/codeforces-go
func getMaxLength(n int, a []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i+2 < n; i++ {
		if a[i] < a[i+1] {
			st := i
			for ; i+2 < n && a[i] < a[i+1]; i++ {
			}
			if a[i] <= a[i+1] {
				continue
			}
			for ; i+1 < n && a[i] > a[i+1]; i++ {
			}
			ans = max(ans, i-st+1)
			i--
		}
	}
	return
}
