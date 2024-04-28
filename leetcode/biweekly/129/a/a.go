package main

// https://space.bilibili.com/206214
func canMakeSquare(grid [][]byte) bool {
	check := func(i, j int) bool {
		cnt := [2]int{}
		cnt[grid[i][j]&1]++
		cnt[grid[i][j+1]&1]++
		cnt[grid[i+1][j]&1]++
		cnt[grid[i+1][j+1]&1]++
		return cnt[0] >= 3 || cnt[1] >= 3
	}
	return check(0, 0) || check(0, 1) || check(1, 0) || check(1, 1)
}
