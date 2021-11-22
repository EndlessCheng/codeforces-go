package main

/* O(n) 做法

先判断最左右的两栋房子颜色是否相同，若不同直接返回 $n-1$。

若相同，则可以在其余位置（尽量靠近左或右）找到颜色不同于最左右房子颜色的房子，计算其到左右的最远距离，即为答案。
*/

// github.com/EndlessCheng/codeforces-go
func maxDistance(colors []int) (ans int) {
	n := len(colors)
	c := colors[0]
	if c != colors[n-1] { return n - 1 }
	l, r := 1, n-2
	for colors[l] == c { l++ }
	for colors[r] == c { r-- }
	return max(r, n-1-l)
}

func max(a, b int) int { if b > a { return b }; return a}
