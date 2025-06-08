package main

// https://space.bilibili.com/206214
func countPermutations(complexity []int) int {
	const mod = 1_000_000_007
	ans := 1
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
		ans = ans * i % mod
	}
	return ans
}
