package main

import (
	"sort"
)

// https://space.bilibili.com/206214
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	done := make([]int, n)
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		exam, study := n, 0
		for i := mx - 1; i >= 0 && study <= i+1; i-- {
			idx := changeIndices[i] - 1
			if done[idx] != mx {
				done[idx] = mx
				exam-- // 考试
				study += nums[idx] // 需要复习的天数
			} else if study > 0 {
				study-- // 复习
			}
		}
		return exam == 0 && study == 0 // 考完了并且复习完了
	})
	if ans > m {
		return -1
	}
	return ans
}
