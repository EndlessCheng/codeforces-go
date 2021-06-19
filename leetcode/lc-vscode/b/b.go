package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func halfQuestions(a []int) (ans int) {
	cnt := [1001]int{}
	for _, c := range a {
		cnt[c]++
	}
	sort.Ints(cnt[:])
	for i, n := 1000, len(a)/2; n > 0; i-- {
		ans++
		n -= cnt[i]
	}
	return
}
