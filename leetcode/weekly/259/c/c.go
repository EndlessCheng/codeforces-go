package main

// github.com/EndlessCheng/codeforces-go
type DetectSquares struct{}

var row [1001][]int
var col = [1001]map[int]int{}

func Constructor() (_ DetectSquares) {
	row = [1001][]int{}
	for i := range col {
		col[i] = map[int]int{}
	}
	return
}

func (DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	row[y] = append(row[y], x)
	col[x][y]++
}

func (DetectSquares) Count(point []int) (ans int) {
	x, y := point[0], point[1]
	for _, x2 := range row[y] {
		if x2 != x {
			d := abs(x2 - x) // 横向距离
			ans += col[x][y-d] * col[x2][y-d] // 将 x-x2 当作正方形顶边两个点，计算正方形底边两点的组合方案
			ans += col[x][y+d] * col[x2][y+d] // 将 x-x2 当作正方形底边两个点，计算正方形顶边两点的组合方案
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
