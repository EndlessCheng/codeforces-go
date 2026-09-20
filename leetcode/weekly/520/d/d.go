package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func largestPower1(nums []int) []int {
	ans := [15]int{}
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	maxWidth := bits.Len(uint(nums[0]))
	for i := maxWidth - 1; i >= 0; i-- {
		// 找最长前缀连续 1
		j := 0
		for j < len(nums) && nums[j]>>i&1 > 0 {
			j++
		}
		ans[14-i] = j

		// [0, j-1] 这一位都是 1，后面无关紧要，为方便排序，全置为 0
		for ; j < len(nums); j++ {
			nums[j] &^= 1 << i
		}
		slices.SortFunc(nums, func(a, b int) int { return b - a })
	}
	return ans[:]
}

func largestPower(nums []int) []int {
	ans := [15]int{}
	groups := [][]int{nums}
	maxWidth := bits.Len(uint(slices.Max(nums)))

	for i := maxWidth - 1; i >= 0; i-- {
		nxt := [][]int{}
		cnt := 0

		for j, a := range groups {
			// 手动排 groups[j]
			var ones, zeros []int
			for _, x := range a {
				if x>>i&1 > 0 {
					ones = append(ones, x)
				} else {
					zeros = append(zeros, x)
				}
			}

			cnt += len(ones)
			// 1 排前面，0 排后面
			if len(ones) > 0 {
				nxt = append(nxt, ones)
			}
			if len(zeros) > 0 {
				nxt = append(nxt, zeros)
				// 遇到 0，剩下的 groups[j] 不再排序
				nxt = append(nxt, groups[j+1:]...)
				break
			}
		}

		ans[14-i] = cnt
		groups = nxt
	}

	return ans[:]
}
