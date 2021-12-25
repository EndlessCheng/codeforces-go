package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func mostWordsFound(sentences []string) (ans int) {
	for _, s := range sentences {
		cnt := strings.Count(s, " ") + 1
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
