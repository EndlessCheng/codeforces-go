package main

// https://space.bilibili.com/206214/dynamic
func equalPairs(grid [][]int) (ans int) {
	cnt := map[[200]int]int{}
	for _, row := range grid {
		a := [200]int{}
		for j, v := range row {
			a[j] = v
		}
		cnt[a]++
	}
	for j := range grid[0] {
		a := [200]int{}
		for i, row := range grid {
			a[i] = row[j]
		}
		ans += cnt[a]
	}
	return
}
