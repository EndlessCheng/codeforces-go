package main

// github.com/EndlessCheng/codeforces-go
func minCost(startPos, homePos, rowCosts, colCosts []int) int {
	x0, y0 := startPos[0], startPos[1]
	x1, y1 := homePos[0], homePos[1]

	// 起点的代价不计入，先减去
	ans := -rowCosts[x0] - colCosts[y0]

	// 累加代价（包含起点）
	for _, cost := range rowCosts[min(x0, x1) : max(x0, x1)+1] {
		ans += cost
	}
	for _, cost := range colCosts[min(y0, y1) : max(y0, y1)+1] {
		ans += cost
	}

	return ans
}
