package main

// github.com/EndlessCheng/codeforces-go
func threeConsecutiveOdds(a []int) (ans bool) {
	for i := 2; i < len(a); i++ {
		if a[i-2]&1+a[i-1]&1+a[i]&1 == 3 {
			return true
		}
	}
	return
}
