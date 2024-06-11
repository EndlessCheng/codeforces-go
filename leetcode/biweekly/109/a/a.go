package main

// https://space.bilibili.com/206214
func isGood(nums []int) bool {
	n := len(nums) - 1
	cnt := make([]int, n+1)
	for _, v := range nums {
		if v > n || v == n && cnt[v] > 1 || v < n && cnt[v] > 0 {
			return false
		}
		cnt[v]++
	}
	return true
}
