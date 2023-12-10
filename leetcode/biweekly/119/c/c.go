package main

// https://space.bilibili.com/206214
func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, x := range nums {
		cnt[x]++
		for cnt[x] > k {
			cnt[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
