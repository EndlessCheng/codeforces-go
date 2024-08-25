package main

import (
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
func countPairs(nums []int) (ans int) {
	slices.Sort(nums)
	cnt := map[int]int{}
	for _, x := range nums {
		set := map[int]struct{}{x: {}} // 不交换
		s := []byte(strconv.Itoa(x))
		m := len(s)
		for i := range s {
			for j := i + 1; j < m; j++ {
				s[i], s[j] = s[j], s[i]
				set[atoi(s)] = struct{}{} // 交换一次
				s[i], s[j] = s[j], s[i]
			}
		}
		for x := range set {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return
}

// 手动转 int 快一些
func atoi(s []byte) (res int) {
	for _, b := range s {
		res = res*10 + int(b&15)
	}
	return
}
