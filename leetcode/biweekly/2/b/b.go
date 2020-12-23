package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func highFive(items [][]int) (ans [][]int) {
	mp := map[int][]int{}
	for _, p := range items {
		mp[p[0]] = append(mp[p[0]], p[1])
	}
	for id, a := range mp {
		sort.Ints(a)
		sum := 0
		for _, v := range a[len(a)-5:] {
			sum += v
		}
		ans = append(ans, []int{id, sum / 5})
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i][0] < ans[j][0] })
	return
}
