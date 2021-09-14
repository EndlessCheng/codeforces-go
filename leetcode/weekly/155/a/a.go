package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimumAbsDifference(a []int) (ans [][]int) {
	sort.Ints(a)
	mi := int(1e9)
	for i := 1; i < len(a); i++ {
		if d := a[i] - a[i-1]; d < mi {
			mi = d
			ans = [][]int{{a[i-1], a[i]}}
		} else if d == mi {
			ans = append(ans, []int{a[i-1], a[i]})
		}
	}
	return
}
