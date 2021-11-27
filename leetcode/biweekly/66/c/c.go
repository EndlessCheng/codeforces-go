package main

// github.com/EndlessCheng/codeforces-go
func minCost(startPos, homePos, rowCosts, colCosts []int) int {
	x0, y0, x1, y1 := startPos[0], startPos[1], homePos[0], homePos[1]
	ans := -rowCosts[x0] - colCosts[y0] // 初始的行列无需计算
	if x0 > x1 { x0, x1 = x1, x0 } // 交换位置，保证 x0 <= x1
	if y0 > y1 { y0, y1 = y1, y0 } // 交换位置，保证 y0 <= y1
	for _, cost := range rowCosts[x0 : x1+1] { ans += cost } // 统计答案
	for _, cost := range colCosts[y0 : y1+1] { ans += cost } // 统计答案
	return ans
}
