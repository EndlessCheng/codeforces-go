package main

// https://space.bilibili.com/206214
func countDistinctIntegers(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
		rev := 0
		for ; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		set[rev] = struct{}{}
	}
	return len(set)
}
