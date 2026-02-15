package main

// https://space.bilibili.com/206214
func rob1(nums, colors []int) int64 {
	n := len(nums)
	f := make([]int, n+1)
	f[1] = nums[0]
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			f[i+1] = f[i] + nums[i]
		} else {
			f[i+1] = max(f[i-1]+nums[i], f[i]) // 选或不选
		}
	}
	return int64(f[n])
}

func rob(nums, colors []int) int64 {
	n := len(nums)
	f0, f1 := 0, nums[0]
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			f0 = f1
			f1 += nums[i]
		} else {
			f0, f1 = f1, max(f0+nums[i], f1)
		}
	}
	return int64(f1)
}
