package main

import "unsafe"

// https://space.bilibili.com/206214
func mergeAdjacent(nums []int) []int64 {
	st := nums[:0] // 原地
	for _, x := range nums {
		st = append(st, x)
		for len(st) > 1 && st[len(st)-1] == st[len(st)-2] {
			st = st[:len(st)-1]
			st[len(st)-1] *= 2
		}
	}
	// 力扣的 int 就是 int64，直接 O(1) 转成 []int64
	return *(*[]int64)(unsafe.Pointer(&st))
}
