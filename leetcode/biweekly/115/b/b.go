package main

// https://space.bilibili.com/206214
func getWordsInLongestSubsequence(words []string, groups []int) (ans []string) {
	n := len(groups)
	for i, x := range groups {
		if i == n-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return
}
