package main

// https://space.bilibili.com/206214
func minimumLevels(possible []int) int {
	// cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
	n := len(possible)
	s := 0
	for _, x := range possible {
		s += x
	}
	s = s*2 - n
	pre := 0
	for i, x := range possible[:n-1] {
		pre += x*4 - 2
		if pre > s {
			return i + 1
		}
	}
	return -1
}
