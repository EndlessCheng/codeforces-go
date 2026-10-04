package main

// https://space.bilibili.com/206214
func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1
	for i < len(nums) {
		if nums[i]%2 == 0 {
			i += 2 // 寻找偶数下标中最左边的奇数
		} else if nums[j]%2 == 1 {
			j += 2 // 寻找奇数下标中最左边的偶数
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i += 2
			j += 2
		}
	}
	return nums
}
