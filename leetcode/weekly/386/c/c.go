package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	lastT := make([]int, n)
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		clear(lastT)
		for t, idx := range changeIndices[:mx] {
			lastT[idx-1] = t + 1
		}
		if slices.Contains(lastT, 0) { // 有课程没有考试时间
			return false
		}

		cnt := 0
		for i, idx := range changeIndices[:mx] {
			idx--
			if i == lastT[idx]-1 { // 考试
				if nums[idx] > cnt { // 没时间复习
					return false
				}
				cnt -= nums[idx] // 复习这门课程
			} else {
				cnt++ // 留着后面用
			}
		}
		return true
	})
	if ans > m {
		return -1
	}
	return ans
}