package main

import "sort"

func isPossibleDivide(nums []int, k int) (ans bool) {
	if len(nums)%k != 0 {
		return
	}

	sort.Ints(nums)
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	consume := func(l, r int) bool {
		for i := l; i < r; i++ {
			if cnt[i] == 0 {
				return false
			}
			cnt[i]--
		}
		return true
	}
	for _, v := range nums {
		if cnt[v] > 0 && !consume(v, v+k) {
			return
		}
	}
	return true
}
