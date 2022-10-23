package main

import "sort"

// https://space.bilibili.com/206214
func makeSimilar(nums, target []int) (ans int64) {
	sort.Ints(nums)
	sort.Ints(target)
	j := [2]int{}
	for _, x := range nums {
		p := x % 2
		for target[j[p]]%2 != p {
			j[p]++
		}
		ans += int64(abs(x - target[j[p]]))
		j[p]++
	}
	return ans / 4
}

func abs(x int) int { if x < 0 { return -x }; return x }
