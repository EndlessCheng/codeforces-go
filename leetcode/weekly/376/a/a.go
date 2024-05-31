package main

// https://space.bilibili.com/206214


func findMissingAndRepeatedValues2(grid [][]int) []int {
	n := len(grid)
	cnt := make([]int, n*n+1)
	for _, row := range grid {
		for _, x := range row {
			cnt[x]++
		}
	}
	ans := [2]int{}
	for i := 1; i <= n*n; i++ {
		if cnt[i] == 2 {
			ans[0] = i
		} else if cnt[i] == 0 {
			ans[1] = i
		}
	}
	return ans[:]
}
