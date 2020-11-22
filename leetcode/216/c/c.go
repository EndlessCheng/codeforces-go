package main

// github.com/EndlessCheng/codeforces-go
func waysToMakeFair(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+2)
	for i, v := range a {
		sum[i+2] = sum[i] + v
	}
	for i := range a {
		if i&1 == 1 {
			s2 := sum[i] + sum[n+1-(n&1^1)] - sum[i+1]
			s1 := sum[i+1] + sum[n+1-(n&1)] - sum[i+2]
			if s1 == s2 {
				ans++
			}
		} else {
			s1 := sum[i] + sum[n+1-(n&1)] - sum[i+1]
			s2 := sum[i+1] + sum[n+1-(n&1^1)] - sum[i+2]
			if s1 == s2 {
				ans++
			}
		}
	}
	return
}
