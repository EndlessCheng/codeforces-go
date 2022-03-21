package main

// github.com/EndlessCheng/codeforces-go
func divideArray(nums []int) bool {
	parity := map[int]int{}
	odd := 0
	for _, v := range nums {
		parity[v] ^= 1
		odd += parity[v]<<1 - 1
	}
	return odd == 0
}
