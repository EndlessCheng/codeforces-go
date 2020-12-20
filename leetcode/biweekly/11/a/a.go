package main

// github.com/EndlessCheng/codeforces-go
func missingNumber(a []int) (ans int) {
	n := len(a)
	d := (a[n-1] - a[0]) / n
	// 大意了
	if d == 0 {
		return a[0]
	}
	for i := 1; ; i++ {
		if a[i]-a[i-1] != d {
			return a[i-1] + d
		}
	}
}
