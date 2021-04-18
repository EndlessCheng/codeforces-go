package main

// github.com/EndlessCheng/codeforces-go
func minOperations(a []int) (ans int) {
	cur := a[0] + 1
	for _, v := range a[1:] {
		if v < cur {
			ans += cur - v
			cur++
		} else {
			cur = v + 1
		}
	}
	return
}
