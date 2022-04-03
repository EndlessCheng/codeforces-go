package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func findWinners(matches [][]int) [][]int {
	man := map[int]struct{}{}
	cnt := map[int]int{}
	for _, p := range matches {
		man[p[0]] = struct{}{}
		man[p[1]] = struct{}{} // 记录所有出现过的人
		cnt[p[1]]++            // 统计每个人输的次数
	}
	ans := [][]int{{}, {}}
	for v := range man {
		if c := cnt[v]; c < 2 {
			ans[c] = append(ans[c], v)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
