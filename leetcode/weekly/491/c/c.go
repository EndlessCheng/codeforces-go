package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func minimumOR(grid [][]int) (ans int) {
	or := 0
	for _, row := range grid {
		// 每行选个最小值，计算 OR
		or |= slices.Min(row)
	}
	// 答案 <= or，那么答案的二进制长度也 <= or 的二进制长度

	// 试填法：ans 的第 i 位能不能是 0？
	// 如果在每一行的能选的数字中，都存在第 i 位是 0 的数，那么 ans 的第 i 位可以是 0，否则必须是 1
	for i := bits.Len(uint(or)) - 1; i >= 0; i-- {
		mask := ans | (1<<i - 1) // mask 低于 i 的比特位全是 1
	next:
		for _, row := range grid {
			for _, x := range row {
				// x 的高于 i 的比特位，如果 mask 是 0，那么 x 的这一位必须也是 0（注意 mask 继承了 ans 高位中的 0）
				// x 的低于 i 的比特位，随意
				// x 的第 i 个比特位，我们期望它是 0
				if x|mask == mask { // x 可以选，且第 i 位是 0
					continue next
				}
			}
			// 这一行的可选数字中，第 i 位全是 1
			ans |= 1 << i // ans 第 i 位必须是 1
			break // 填下一位
		}
	}
	return
}
