package main

import "strings"

// https://space.bilibili.com/206214
func bestClosingTime(customers string) int {
	penalty := strings.Count(customers, "Y")
	minPenalty := penalty
	ans := 0 // [0,n-1] 是第二段
	for i, c := range customers {
		if c == 'N' {
			penalty++
		} else {
			penalty--
		}
		if penalty < minPenalty {
			minPenalty = penalty
			ans = i + 1 // [0,i] 是第一段，[i+1,n-1] 是第二段
		}
	}
	return ans
}
