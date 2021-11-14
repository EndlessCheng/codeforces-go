package main

// github.com/EndlessCheng/codeforces-go
func checkAlmostEquivalent(word1, word2 string) bool {
	diff := [26]int{}
	for _, ch := range word1 {
		diff[ch-'a']++
	}
	for _, ch := range word2 {
		diff[ch-'a']--
	}
	for _, d := range diff {
		if d < -3 || d > 3 {
			return false
		}
	}
	return true
}
