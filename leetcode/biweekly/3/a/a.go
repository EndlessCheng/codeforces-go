package main

// github.com/EndlessCheng/codeforces-go
func twoSumLessThanK(a []int, k int) (ans int) {
	ans = -1
	for i, v := range a {
		for j := i + 1; j < len(a); j++ {
			w := a[j]
			if v+w < k && v+w > ans {
				ans = v + w
			}
		}
	}
	return
}
