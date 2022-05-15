package main

// github.com/EndlessCheng/codeforces-go
func removeAnagrams(words []string) []string {
	ans := []string{words[0]}
	for _, word := range words[1:] {
		cnt := [26]int{}
		for _, b := range word {
			cnt[b-'a']++
		}
		for _, b := range ans[len(ans)-1] {
			cnt[b-'a']--
		}
		if cnt != [26]int{} {
			ans = append(ans, word)
		}
	}
	return ans
}
