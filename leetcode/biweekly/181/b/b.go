package main

// https://space.bilibili.com/206214
func compareBitonicSums(nums []int) int {
	diff := 0
	inc := true
	for i, x := range nums {
		if i > 0 && nums[i-1] < x && x > nums[i+1] {
			inc = false
			// 注意峰顶抵消掉了，不算入 diff
		} else if inc {
			diff += x
		} else {
			diff -= x
		}
	}

	if diff > 0 {
		return 0
	}
	if diff < 0 {
		return 1
	}
	return -1
}
