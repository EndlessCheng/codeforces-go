package main

// https://space.bilibili.com/206214
func findValidSplit(nums []int) int {
	n := len(nums)
	left := map[int]int{}   // left[p] 表示质数 p 首次出现的下标
	right := make([]int, n) // right[i] 表示左端点为 i 的区间的右端点的最大值
	for i := range right {
		right[i] = -1
	}
	f := func(p, i int) {
		if l, ok := left[p]; ok {
			right[l] = i // 记录左端点 l 对应的右端点的最大值
		} else {
			left[p] = i // 第一次遇到质数 p
		}
	}
	for i, x := range nums {
		for d := 2; d*d <= x; d++ { // 分解质因子
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	maxR := right[0]
	for l, r := range right {
		if l > maxR { // 最远可以到 maxR
			return maxR
		}
		maxR = max(maxR, r)
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
