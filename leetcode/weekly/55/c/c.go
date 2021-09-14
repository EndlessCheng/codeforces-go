package main

// github.com/EndlessCheng/codeforces-go
func numSubarrayProductLessThanK(a []int, k int) (ans int) {
	pre := make([]int, len(a))
	p := -1
	for i, v := range a {
		pre[i] = p
		if v > 1 {
			p = i
		}
	}
	for i := range a {
		st := i
		for v := 1; i >= 0 && v*a[i] < k; i = pre[i] {
			v *= a[i]
		}
		ans += st - i
	}
	return
}
