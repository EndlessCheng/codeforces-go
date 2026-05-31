package main

// https://space.bilibili.com/206214
func maximumMEX(nums []int) (ans []int) {
	n := len(nums)
	// mex 最大是 n，>= n 的数无需考虑
	pos := make([][]int, n+1) // n 作为哨兵
	for i, x := range nums {
		if x < n {
			pos[x] = append(pos[x], i)
		}
	}

	for i := 0; i < n; i++ {
		start := i // 这一段的起点
		// 枚举这一段的 mex，越大越好（字典序越大）
		mex := 0
		for ; ; mex++ {
			// 清理在 start 之前的下标
			for len(pos[mex]) > 0 && pos[mex][0] < start {
				pos[mex] = pos[mex][1:]
			}
			if len(pos[mex]) == 0 {
				break
			}
			// 这一段包含剩余元素中的最左边的 mex
			i = max(i, pos[mex][0])
		}
		ans = append(ans, mex)
	}
	return
}
