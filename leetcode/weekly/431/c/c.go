package main

import "slices"

// https://space.bilibili.com/206214
// 2271. 毯子覆盖的最多白色砖块数
func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
	cover, left := 0, 0
	for _, tile := range tiles {
		tl, tr, c := tile[0], tile[1], tile[2]
		cover += (tr - tl + 1) * c
		for tiles[left][1]+carpetLen-1 < tr {
			cover -= (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2]
			left++
		}
		uncover := max((tr-carpetLen+1-tiles[left][0])*tiles[left][2], 0)
		ans = max(ans, cover-uncover)
	}
	return
}

func maximumCoins(coins [][]int, k int) int64 {
	slices.SortFunc(coins, func(a, b []int) int { return a[0] - b[0] })
	ans := maximumWhiteTiles(coins, k)

	slices.Reverse(coins)
	for _, t := range coins {
		t[0], t[1] = -t[1], -t[0]
	}
	return int64(max(ans, maximumWhiteTiles(coins, k)))
}
