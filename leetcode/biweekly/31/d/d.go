package main

func minNumberOperations(target []int) int {
	ans := target[0]
	for i := 1; i < len(target); i++ {
		ans += max(target[i]-target[i-1], 0)
	}
	return ans
}
