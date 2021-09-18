package main

// O(n) 做法

// github.com/EndlessCheng/codeforces-go
func countKDifference(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for _, v := range nums {
		ans += cnt[v-k]
	}
	return
}
