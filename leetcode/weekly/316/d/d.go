package main

import "sort"

// https://space.bilibili.com/206214
func f(a []int) {
	for i, v := range a {
		if v%2 > 0 {
			a[i] = -v // 由于元素都是正数，把奇数变成负数，这样排序后奇偶就自动分开了
		}
	}
	sort.Ints(a)
}

func makeSimilar(nums, target []int) (ans int64) {
	f(nums)
	f(target)
	for i, v := range nums {
		ans += int64(abs(v - target[i]))
	}
	return ans / 4
}

func abs(x int) int { if x < 0 { return -x }; return x }
