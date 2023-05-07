package main

// https://space.bilibili.com/206214
func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suf := make([]int, n+1)
	set := map[int]struct{}{}
	for i := n - 1; i >= 0; i-- {
		set[nums[i]] = struct{}{}
		suf[i] = len(set)
	}

	set = make(map[int]struct{}, len(set))
	ans := make([]int, n)
	for i, x := range nums {
		set[x] = struct{}{}
		ans[i] = len(set) - suf[i+1]
	}
	return ans
}
