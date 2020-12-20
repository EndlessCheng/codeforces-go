package main

import "sort"

type Sea struct {
	points [][]int
}

func (s Sea) hasShips(topRight, bottomLeft []int) bool {
	for _, p := range s.points {
		if bottomLeft[0] <= p[0] && p[0] <= topRight[0] && bottomLeft[1] <= p[1] && p[1] <= topRight[1] {
			return true
		}
	}
	return false
}

func main() {
	s := Sea{[][]int{{6, 6}, {100, 50}, {999, 81}, {50, 50}, {700, 600}}}
	println(countShips(s, []int{1000, 1000}, []int{0, 0}))
}

// github.com/EndlessCheng/codeforces-go
func countShips(s Sea, topRight, bottomLeft []int) (ans int) {
	var f func([]int, []int)
	f = func(bottomLeft, topRight []int) {
		if !s.hasShips(topRight, bottomLeft) {
			return
		}
		ans++
		x := sort.Search(topRight[0], func(x int) bool { return x >= bottomLeft[0] && s.hasShips([]int{x, topRight[1]}, bottomLeft) })
		y := sort.Search(topRight[1], func(y int) bool { return y >= bottomLeft[1] && s.hasShips([]int{x, y}, bottomLeft) })
		if x < topRight[0] {
			f([]int{x + 1, bottomLeft[1]}, topRight)
		}
		if y < topRight[1] {
			f([]int{bottomLeft[0], y + 1}, []int{x, topRight[1]})
		}
	}
	f(bottomLeft, topRight)
	return
}
