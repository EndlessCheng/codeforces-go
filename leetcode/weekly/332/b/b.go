package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func countFairPairs(nums []int, lower, upper int) int64 {
	slices.Sort(nums)
	count := func(upper int) (res int64) {
		i, j := 0, len(nums)-1
		for i < j {
			if nums[i]+nums[j] <= upper {
				res += int64(j - i)
				i++
			} else {
				j--
			}
		}
		return res
	}
	return count(upper) - count(lower-1)
}

func countFairPairs2(nums []int, lower, upper int) (ans int64) {
	slices.Sort(nums)
	left, right := len(nums), len(nums)
	for j, x := range nums {
		for right > 0 && nums[right-1] > upper-x {
			right--
		}
		for left > 0 && nums[left-1] >= lower-x {
			left--
		}
		ans += int64(min(right, j) - min(left, j))
	}
	return
}

func countFairPairs1(nums []int, lower, upper int) (ans int64) {
	slices.Sort(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r - l)
	}
	return
}
