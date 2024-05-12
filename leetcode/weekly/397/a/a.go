package main

// https://space.bilibili.com/206214
func findPermutationDifference(s, t string) (ans int) {
	pos := [26]int{}
	for i, c := range s {
		pos[c-'a'] = i
	}
	for i, c := range t {
		ans += abs(i - pos[c-'a'])
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
