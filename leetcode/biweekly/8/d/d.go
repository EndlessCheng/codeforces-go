package main

import "sort"

// 注：题解区有 O(1) 做法

// github.com/EndlessCheng/codeforces-go
func maximumNumberOfOnes(width, height, sz, maxOnes int) (ans int) {
	cnt := []int{}
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			cnt = append(cnt, ((width-i-1)/sz+1)*((height-j-1)/sz+1))
		}
	}
	sort.Ints(cnt)
	for _, c := range cnt[len(cnt)-maxOnes:] {
		ans += c
	}
	return
}
