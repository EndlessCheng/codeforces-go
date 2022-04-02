package main

// github.com/EndlessCheng/codeforces-go
func triangularSum(a []int) int {
	for n := len(a) - 1; n > 0; n-- {
		for i := 0; i < n; i++ {
			a[i] = (a[i] + a[i+1]) % 10
		}
	}
	return a[0]
}
