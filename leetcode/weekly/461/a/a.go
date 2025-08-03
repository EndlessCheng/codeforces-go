package main

// https://space.bilibili.com/206214
func isTrionic(nums []int) bool {
	n := len(nums)
	// 第一段
	i := 1
	for i < n && nums[i-1] < nums[i] {
		i++
	}
	if i == 1 { // 第一段至少要有两个数
		return false
	}

	// 第二段
	i0 := i
	for i < n && nums[i-1] > nums[i] {
		i++
	}
	if i == i0 || i == n { // 第二段至少要有两个数，第三段至少要有两个数
		return false
	}

	// 第三段
	for i < n && nums[i-1] < nums[i] {
		i++
	}
	return i == n
}
