package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
	slices.SortFunc(tiles, func(a, b []int) int { return a[0] - b[0] })
	cover, left := 0, 0
	for _, tile := range tiles {
		tl, tr := tile[0], tile[1]
		cover += tr - tl + 1

		carpetLeft := tr - carpetLen + 1  // 毯子左端点
		for tiles[left][1] < carpetLeft { // tiles[left] 完全不在窗口中
			cover -= tiles[left][1] - tiles[left][0] + 1
			left++
		}

		// tiles[left] 中的小于 carpetLeft 的这部分不在窗口中
		uncover := max(carpetLeft-tiles[left][0], 0)
		ans = max(ans, cover-uncover)
	}
	return
}
