package main

// https://space.bilibili.com/206214
func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	for i, x := range nums {
		if i > 0 && f[i-1] && x == nums[i-1] ||
			i > 1 && f[i-2] && (x == nums[i-1] && x == nums[i-2] ||
				                x == nums[i-1]+1 && x == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}
