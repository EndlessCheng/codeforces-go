package main

// https://space.bilibili.com/206214
func checkDistances(s string, distance []int) bool {
	last := [26]int{}
	for i, c := range s {
		c -= 'a'
		if last[c] > 0 && i-last[c] != distance[c] {
			return false
		}
		last[c] = i + 1
	}
	return true
}
