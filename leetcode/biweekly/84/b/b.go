package main

// https://space.bilibili.com/206214
func countBadPairs(nums []int) int64 {
	n := len(nums)
	ans := n * (n - 1) / 2
	cnt := map[int]int{}
	for i, x := range nums {
		ans -= cnt[x-i]
		cnt[x-i]++
	}
	return int64(ans)
}
