package main

import "sort"

func smallestDivisor(nums []int, threshold int) int {
	ans := sort.Search(int(1e6+1), func(low int) bool {
		if low == 0 {
			return false
		}
		sum := 0
		for _, v := range nums {
			sum += (v-1)/low + 1
		}
		return sum <= threshold
	})
	return ans
}
