package main

// https://space.bilibili.com/206214/dynamic
func equalPairs(grid [][]int) (ans int) {
	cnt := map[[200]int]int{}
	a := [200]int{}
	for _, row := range grid {
		copy(a[:], row)
		cnt[a]++
	}
	for j := range grid[0] {
		for i, row := range grid {
			a[i] = row[j]
		}
		ans += cnt[a]
	}
	return
}
