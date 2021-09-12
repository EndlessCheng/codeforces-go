package main

/*
哈希表统计比值的最简分数

将宽高比 w/h 化简为最简分数，用哈希表统计最简分数的个数。若有 $m$ 个宽高比相同的矩形，则可互换的矩形对有 $C_m2=\dfrac{m\cdot(m-1}{2}$ 个。遍历哈希表，累加所有可互换的矩形对的个数即为答案。

*/

// github.com/EndlessCheng/codeforces-go
func interchangeableRectangles(rectangles [][]int) (ans int64) {
	cnt := map[[2]int]int64{}
	for _, p := range rectangles {
		// 计算 w/h 的最简分数，计入哈希表
		w, h := p[0], p[1]
		g := gcd(w, h)
		cnt[[2]int{w / g, h / g}]++
	}
	for _, m := range cnt {
		ans += m * (m - 1) / 2
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
