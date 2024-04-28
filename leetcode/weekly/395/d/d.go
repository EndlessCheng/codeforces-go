package main

import "sort"

// https://space.bilibili.com/206214
func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	k := (n*(n+1)/2 + 1) / 2
	ans := 1 + sort.Search(n-1, func(upper int) bool {
		upper++
		cnt := 0
		l := 0
		freq := map[int]int{}
		for r, in := range nums {
			freq[in]++
			for len(freq) > upper {
				out := nums[l]
				freq[out]--
				if freq[out] == 0 {
					delete(freq, out)
				}
				l++
			}
			cnt += r - l + 1
			if cnt >= k {
				return true
			}
		}
		return false
	})
	return ans
}
