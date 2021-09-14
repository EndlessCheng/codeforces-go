package main

import "strings"

/*

直接模拟就行了，或者 AC 自动机

 */

// github.com/EndlessCheng/codeforces-go
func numOfStrings(patterns []string, word string) (ans int) {
	for _, p := range patterns {
		if strings.Contains(word, p) {
			ans++
		}
	}
	return
}
