package main

// https://space.bilibili.com/206214
func minSwaps(nums, forbidden []int) int {
	n := len(nums)
	total := map[int]int{}
	for _, x := range nums {
		total[x]++
	}

	cnt := map[int]int{}
	k, mx := 0, 0
	for i, x := range forbidden {
		total[x]++
		if total[x] > n {
			return -1
		}
		if x == nums[i] {
			k++
			cnt[x]++
			mx = max(mx, cnt[x])
		}
	}

	return max((k+1)/2, mx)
}
