package main

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	sort.Ints(nums)
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	consume := func(st, end int) bool {
		for i := st; i < end; i++ {
			if cnt[i] > 0 {
				cnt[i]--
			} else {
				return false
			}
		}
		return true
	}
	for _, v := range nums {
		if cnt[v] > 0 && !consume(v, v+k) {
			return false
		}
	}
	return true
}
