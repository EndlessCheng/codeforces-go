package main

// github.com/EndlessCheng/codeforces-go
func minimumCardPickup(cards []int) int {
	ans := len(cards) + 1
	pos := map[int]int{}
	for i, v := range cards {
		if p, ok := pos[v]; ok && i-p+1 < ans {
			ans = i - p + 1
		}
		pos[v] = i
	}
	if ans <= len(cards) {
		return ans
	}
	return -1
}
