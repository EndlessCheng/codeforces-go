package main

// https://space.bilibili.com/206214
func rotateElements(nums []int, k int) []int {
	// 取出非负数
	a := []int{}
	for _, x := range nums {
		if x >= 0 {
			a = append(a, x)
		}
	}

	// 双指针，把 a 填入 nums，跳过负数
	j := k
	for i, x := range nums {
		if x >= 0 {
			nums[i] = a[j%len(a)]
			j++
		}
	}
	return nums
}
