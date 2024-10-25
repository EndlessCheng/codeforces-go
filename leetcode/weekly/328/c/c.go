package main

// https://space.bilibili.com/206214
func countGood(nums []int, k int) (ans int64) {
	cnt := map[int]int{}
	pairs, left := 0, 0
	for _, x := range nums {
		pairs += cnt[x]
		cnt[x]++
		for pairs >= k {
			cnt[nums[left]]--
			pairs -= cnt[nums[left]]
			left++
		}
		ans += int64(left)
	}
	return
}
