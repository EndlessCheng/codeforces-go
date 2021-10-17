package main

// O(2^n) 子集或写法

// github.com/EndlessCheng/codeforces-go
func countMaxOrSubsets(nums []int) (ans int) {
	all := 0
	for _, v := range nums {
		all |= v
	}
	// 时间复杂度为 O(1+2+4+...+2^(n-1)) = O(2^n)
	or := make([]int, 1<<len(nums))
	for i, v := range nums {
		for j, k := 0, 1<<i; j < k; j++ {
			res := or[j] | v
			or[k|j] = res
			if res == all {
				ans++
			}
		}
	}
	return
}
