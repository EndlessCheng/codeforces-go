package main

import "slices"

// https://space.bilibili.com/206214
func relocateMarbles(nums, moveFrom, moveTo []int) []int {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}

	for i, x := range moveFrom {
		delete(set, x)
		set[moveTo[i]] = struct{}{}
	}

	ans := make([]int, 0, len(set))
	for x := range set {
		ans = append(ans, x)
	}
	slices.Sort(ans)
	return ans
}
