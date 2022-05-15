package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })
	cover, left := 0, 0
	for _, t := range tiles {
		tl, tr := t[0], t[1]
		cover += tr - tl + 1
		for tiles[left][1]+carpetLen-1 < tr {
			cover -= tiles[left][1] - tiles[left][0] + 1
			left++
		}
		ans = max(ans, cover-max(tr-carpetLen+1-tiles[left][0], 0)) // 0 表示毯子左端点不在瓷砖内的情况
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
