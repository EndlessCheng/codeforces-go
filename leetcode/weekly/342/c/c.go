package main

// https://space.bilibili.com/206214
func getSubarrayBeauty(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] {
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++
		left := x
		for j, c := range cnt[:bias] {
			left -= c
			if left <= 0 {
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]--
	}
	return ans
}

func getSubarrayBeauty2(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias]int{}
	neg := 0
	for _, num := range nums[:k-1] {
		if num < 0 {
			cnt[num+bias]++
			neg++
		}
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		if num < 0 {
			cnt[num+bias]++
			neg++
		}
		if neg >= x {
			left := x
			for j, c := range cnt {
				left -= c
				if left <= 0 {
					ans[i] = j - bias
					break
				}
			}
		}
		if nums[i] < 0 {
			cnt[nums[i]+bias]--
			neg--
		}
	}
	return ans
}
