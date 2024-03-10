package main

import "math"

// https://space.bilibili.com/206214
func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	f := make([]int, n+1)
	for i := 1; i <= k; i++ {
		pre := f[i-1]
		f[i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		for j := i; j <= n-k+i; j++ {
			mx = max(mx, pre-s[j-1]*w)
			pre = f[j]
			f[j] = max(f[j-1], s[j]*w+mx)
		}
	}
	return int64(f[n])
}

func maximumStrength2(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	f := make([][]int, k+1)
	f[0] = make([]int, n+1)
	for i := 1; i <= k; i++ {
		f[i] = make([]int, n+1)
		f[i][i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		for j := i; j <= n-k+i; j++ { // j 的上下界是因为其它子数组要至少选一个数
			mx = max(mx, f[i-1][j-1]-s[j-1]*w)
			f[i][j] = max(f[i][j-1], s[j]*w+mx)
		}
	}
	return int64(f[k][n])
}
