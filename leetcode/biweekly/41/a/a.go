package main

// github.com/EndlessCheng/codeforces-go
func countConsistentStrings(allowed string, words []string) (ans int) {
	cnt := ['z' + 1]bool{}
	for _, b := range allowed {
		cnt[b] = true
	}
o:
	for _, s := range words {
		for _, b := range s {
			if !cnt[b] {
				continue o
			}
		}
		ans++
	}
	return
}
