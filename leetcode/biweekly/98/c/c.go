package main

// https://space.bilibili.com/206214
func minImpossibleOR(nums []int) int {
	mask := 0
	for _, x := range nums {
		if x&(x-1) == 0 { // x 是 2 的幂次
			mask |= x
		}
	}
	// 取反后，用 lowbit 找第一个 0 比特位
	mask = ^mask
	return mask & -mask
}
