package main

import "sort"

// https://space.bilibili.com/206214
func minTime(s string, order []int, k int) int {
	n := len(s)
	if n*(n+1)/2 < k { // 全改成星号也无法满足要求
		return -1
	}

	star := make([]int, n) // 避免在二分内部反复创建/初始化列表
	ans := sort.Search(len(order), func(m int) bool {
		m++
		for _, j := range order[:m] {
			star[j] = m
		}
		cnt := 0
		last := -1 // 上一个 '*' 的位置
		for i, x := range star {
			if x == m { // s[i] 是 '*'
				last = i
			}
			cnt += last + 1
			if cnt >= k { // 提前退出循环
				return true
			}
		}
		return false
	})
	return ans
}
