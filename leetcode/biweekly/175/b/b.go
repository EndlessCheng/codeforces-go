package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minimumK(nums []int) int {
	n := len(nums)
	left := int(math.Ceil(math.Sqrt(float64(n)))) // 答案的下界
	right := slices.Max(nums)
	ans := left + sort.Search(right-left, func(k int) bool {
		k += left
		sum := n
		for _, x := range nums {
			sum += (x - 1) / k
		}
		return sum <= k*k
	})
	return ans
}
