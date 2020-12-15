package main

// github.com/EndlessCheng/codeforces-go
func numOfSubarrays(a []int, k, threshold int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for r := k; r <= n; r++ {
		if sum[r]-sum[r-k] >= threshold*k {
			ans++
		}
	}
	return
}
