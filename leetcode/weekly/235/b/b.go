package main

// github.com/EndlessCheng/codeforces-go
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	mp := map[int]map[int]struct{}{}
	for _, p := range logs {
		id := p[0]
		if mp[id] == nil {
			mp[id] = map[int]struct{}{}
		}
		mp[id][p[1]] = struct{}{}
	}
	ans := make([]int, k+1)
	for _, m := range mp {
		ans[len(m)]++
	}
	return ans[1:]
}
