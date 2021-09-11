package main

/*
用哈希表统计差异

*/

// github.com/EndlessCheng/codeforces-go
func minimumSwitchingTimes(source, target [][]int) (ans int) {
	cnt := map[int]int{}
	for i, row := range source {
		for j, v := range row {
			cnt[v]++
			cnt[target[i][j]]--
		}
	}
	for _, c := range cnt {
		ans += abs(c)
	}
	return ans / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
