package main

// https://space.bilibili.com/206214
func largestSquareArea(bottomLeft, topRight [][]int) int64 {
	maxSide := 0
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		if t1[0]-b1[0] <= maxSide || t1[1]-b1[1] <= maxSide {
			continue // 最优性剪枝：maxSide 不可能变大
		}
		for j, b2 := range bottomLeft[:i] {
			t2 := topRight[j]
			width := min(t1[0], t2[0]) - max(b1[0], b2[0])  // 右上横坐标 - 左下横坐标
			height := min(t1[1], t2[1]) - max(b1[1], b2[1]) // 右上纵坐标 - 左下纵坐标
			side := min(width, height)
			maxSide = max(maxSide, side)
		}
	}
	return int64(maxSide) * int64(maxSide)
}
