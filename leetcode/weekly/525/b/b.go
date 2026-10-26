package main

import (
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
func maximizeSum(nums []int, k int) (ans int64) {
	for i, x := range nums {
		ans += int64(x)

		// 把 x 从大到小排列，得到 y
		s := []byte(strconv.Itoa(x))
		slices.SortFunc(s, func(a, b byte) int { return int(b) - int(a) })
		y, _ := strconv.Atoi(string(s))

		// 把修改后的增量保存到 nums[i] 中
		nums[i] = y - x
	}

	slices.Sort(nums)
	// 选择增量最大的 k 个数
	for _, d := range nums[len(nums)-k:] {
		ans += int64(d)
	}
	return
}
