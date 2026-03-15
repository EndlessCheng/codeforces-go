package main

// https://space.bilibili.com/206214
func firstUniqueEven(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		if x%2 == 0 {
			cnt[x]++
		}
	}

	for _, x := range nums {
		if x%2 == 0 && cnt[x] == 1 {
			return x
		}
	}
	return -1
}
