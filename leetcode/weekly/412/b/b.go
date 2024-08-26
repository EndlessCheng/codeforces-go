package main

import (
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
var pow10 = [...]int{1, 10, 100, 1000, 10000, 100000, 1000000}

func countPairs(nums []int) int {
	slices.Sort(nums)
	ans := 0
	cnt := make(map[int]int)
	a := [len(pow10)]int{}
	for _, x := range nums {
		st := map[int]struct{}{x: {}} // 不交换
		m := 0
		for v := x; v > 0; v /= 10 {
			a[m] = v % 10
			m++
		}
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				st[x+(a[j]-a[i])*(pow10[i]-pow10[j])] = struct{}{} // 交换一次
			}
		}
		for x := range st {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return ans
}

func countPairs2(nums []int) (ans int) {
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
