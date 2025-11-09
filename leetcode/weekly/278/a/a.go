package main

// github.com/EndlessCheng/codeforces-go
func findFinalValue1(nums []int, original int) int {
	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}
	for has[original] {
		original *= 2
	}
	return original
}

func findFinalValue2(nums []int, original int) int {
	has := map[int]bool{}
	for _, x := range nums {
		k := x / original
		if x%original == 0 && k&(k-1) == 0 {
			has[x] = true
		}
	}
	for has[original] {
		original *= 2
	}
	return original
}

func findFinalValue(nums []int, original int) int {
	mask := 0
	for _, x := range nums {
		k := x / original
		if x%original == 0 && k&(k-1) == 0 {
			mask |= k
		}
	}
	// 找最低的 0，等价于取反后，找最低的 1（lowbit）
	mask = ^mask
	return original * (mask & -mask)
}
