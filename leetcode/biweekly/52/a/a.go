package main

import (
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func sortSentence(s string) (ans string) {
	a := strings.Split(s, " ")
	sort.Slice(a, func(i, j int) bool { return a[i][len(a[i])-1] < a[j][len(a[j])-1] })
	for i, s := range a {
		a[i] = s[:len(s)-1]
	}
	return strings.Join(a, " ")
}
