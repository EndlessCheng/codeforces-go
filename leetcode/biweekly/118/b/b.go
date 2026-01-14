package main

// https://space.bilibili.com/206214
// 128. 最长连续序列
func longestConsecutive(nums []int) (ans int) {
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true // 把 nums 转成哈希集合
	}

	for x := range has { // 遍历哈希集合
		if has[x-1] { // 如果 x 不是序列的起点，直接跳过
			continue
		}
		// x 是序列的起点
		y := x + 1
		for has[y] { // 不断查找下一个数是否在哈希集合中
			y++
		}
		// 循环结束后，y-1 是最后一个在哈希集合中的数
		ans = max(ans, y-x) // 从 x 到 y-1 一共 y-x 个数
	}
	return
}

func maximizeSquareHoleArea(_, _ int, hBars, vBars []int) int {
	side := min(longestConsecutive(hBars), longestConsecutive(vBars)) + 1
	return side * side
}
