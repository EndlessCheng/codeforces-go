package main

// https://space.bilibili.com/206214
func uniqueOccurrences(a []int) bool {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	has := map[int]bool{}
	for _, c := range cnt {
		if has[c] {
			return false
		}
		has[c] = true
	}
	return true
}
