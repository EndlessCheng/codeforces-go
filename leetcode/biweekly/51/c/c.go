package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maximumElementAfterDecrementingAndRearranging1(arr []int) int {
	slices.Sort(arr)
	n := len(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = min(arr[i], arr[i-1]+1)
	}
	return arr[n-1]
}

func maximumElementAfterDecrementingAndRearranging(arr []int) (ans int) {
	n := len(arr)
	cnt := make([]int, n+1)
	for _, x := range arr {
		cnt[min(x, n)]++
	}

	for x := 1; x <= n; x++ {
		ans = min(ans+cnt[x], x)
	}
	return
}
