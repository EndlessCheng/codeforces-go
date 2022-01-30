package main

// github.com/EndlessCheng/codeforces-go
func findFinalValue(nums []int, original int) int {
	maxK := 1
	for _, num := range nums {
		if num%original == 0 {
			k := num / original
			if k&(k-1) == 0 && k > maxK { // 倍数是 2 的幂次
				maxK = k
			}
		}
	}
	return original * maxK
}
