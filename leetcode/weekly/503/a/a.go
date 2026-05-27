package main

// https://space.bilibili.com/206214
func limitOccurrences(nums []int, k int) []int {
	stackSize := k // 栈的大小，前 k 个元素默认保留
	for i := k; i < len(nums); i++ {
		if nums[i] != nums[stackSize-k] { // 和栈的倒数第 k 个数比较
			nums[stackSize] = nums[i] // 入栈
			stackSize++
		}
	}
	return nums[:stackSize]
}
