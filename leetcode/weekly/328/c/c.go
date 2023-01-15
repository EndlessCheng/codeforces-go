package main

// https://space.bilibili.com/206214
func countGood(nums []int, k int) (ans int64) {
	cnt := map[int]int{}
	left, pairs := 0, 0
	for _, x := range nums {
		pairs += cnt[x]
		cnt[x]++
		ans += int64(left)
		for pairs >= k {
			ans++
			cnt[nums[left]]--
			pairs -= cnt[nums[left]]
			left++
		}
	}
	return
}
