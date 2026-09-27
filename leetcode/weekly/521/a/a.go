package main

import "slices"

// https://space.bilibili.com/206214
func rearrangeArray1(nums []int) []int {
	mx := slices.Max(nums)
	cnt := make([]int, mx+1)
	for _, x := range nums {
		cnt[x]++
	}

	n := len(nums)
	ans := make([]int, 0, n)
	for len(ans) < n {
		for x, c := range cnt {
			if c > 0 {
				ans = append(ans, x)
				cnt[x]--
			}
		}
	}
	return ans
}

func rearrangeArray(nums []int) []int {
	n := len(nums)
	mx := slices.Max(nums)
	cnt := make([]int, mx+1)
	levels := make([][]int, n)
	for _, x := range nums {
		c := cnt[x]
		levels[c] = append(levels[c], x)
		cnt[x]++
	}

	ans := make([]int, 0, n)
	// 从下到上遍历每一层的积木
	for _, level := range levels {
		slices.Sort(level)
		ans = append(ans, level...)
	}
	return ans
}
