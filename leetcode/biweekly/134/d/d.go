package main

import "sort"

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k int) (ans int64) {
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		a := nums[:i+1]
		ans += int64(sort.SearchInts(a, k+1) - sort.SearchInts(a, k))
	}
	return
}

func countSubarrays2(nums []int, k int) (ans int64) {
	left, right := 0, 0
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		for left <= i && nums[left] < k {
			left++
		}
		for right <= i && nums[right] <= k {
			right++
		}
		ans += int64(right - left)
	}
	return
}

func countSubarrays3(nums []int, k int) (ans int64) {
	cnt := 0
	for i, x := range nums {
		if x == k {
			cnt++
		}
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			if nums[j] == k {
				cnt--
			}
			nums[j] &= x
			if nums[j] == k {
				cnt++
			}
		}
		ans += int64(cnt)
	}
	return
}
