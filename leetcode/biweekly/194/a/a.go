package main

// https://space.bilibili.com/206214
func peakValleyDifference(nums []int) (ans int) {
	for i := 1; i < len(nums)-1; i++ {
		x, y, z := nums[i-1], nums[i], nums[i+1]
		if x < y && y > z {
			ans += y
		} else if x > y && y < z {
			ans -= y
		}
	}
	return
}
