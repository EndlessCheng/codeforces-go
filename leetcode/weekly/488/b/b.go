package main

import "unsafe"

// https://space.bilibili.com/206214
func mergeAdjacent(nums []int) []int64 {
	st := nums[:0] // 原地
	for _, x := range nums {
		for len(st) > 0 && st[len(st)-1] == x {
			st = st[:len(st)-1]
			x *= 2
		}
		st = append(st, x)
	}
	// 力扣的 int 就是 int64，直接 O(1) 转成 []int64
	return *(*[]int64)(unsafe.Pointer(&st))
}
