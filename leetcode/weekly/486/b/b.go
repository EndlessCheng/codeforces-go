package main

import "slices"

// https://space.bilibili.com/206214
func rotateLeft(a []int, k int) {
	slices.Reverse(a[:k])
	slices.Reverse(a[k:])
	slices.Reverse(a)
}

func rotateElements(nums []int, k int) (ans []int) {
	// 取出非负数
	a := []int{}
	for _, x := range nums {
		if x >= 0 {
			a = append(a, x)
		}
	}

	m := len(a)
	// 没有非负数，无需操作
	if m == 0 {
		return nums
	}

	// 向左轮替 k 个位置
	rotateLeft(a, k%m)

	// 双指针，把 a 填入 nums，跳过负数
	j := 0
	for i, x := range nums {
		if x >= 0 {
			nums[i] = a[j]
			j++
		}
	}
	return nums
}
