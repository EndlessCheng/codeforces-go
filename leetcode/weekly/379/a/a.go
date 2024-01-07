package main

// https://space.bilibili.com/206214
func areaOfMaxDiagonal(dimensions [][]int) (ans int) {
	maxL := 0
	for _, d := range dimensions {
		x, y := d[0], d[1]
		l := x*x + y*y
		if l > maxL || l == maxL && x*y > ans {
			maxL = l
			ans = x * y
		}
	}
	return
}
