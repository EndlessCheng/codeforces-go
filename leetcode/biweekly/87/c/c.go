package main

// https://space.bilibili.com/206214
func smallestSubarrays1(nums []int) []int {
	ans := make([]int, len(nums))
	for i, x := range nums { // 计算右端点为 i 的子数组的或值
		ans[i] = 1 // 子数组的长度至少是 1
		// 循环直到 nums[j] 无法增大，其左侧元素也无法增大
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x       // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
			ans[j] = i - j + 1 // nums[j] 最后一次增大时的子数组长度就是答案
		}
	}
	return ans
}

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[n-1] = 1
	if n == 1 {
		return ans
	}

	// 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
	nums[n-1] |= nums[n-2]
	leftOr, right, bottom := 0, n-1, n-2
	for left := n - 2; left >= 0; left-- {
		leftOr |= nums[left]
		// 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
		for right > left && leftOr|nums[right] == leftOr|nums[right-1] {
			right--
			// 栈中只剩一个数
			if bottom >= right {
				// 重新构建一个栈，栈底为 left，栈顶为 right
				for i := left + 1; i <= right; i++ {
					nums[i] |= nums[i-1]
				}
				bottom = left
				leftOr = 0
			}
		}
		ans[left] = right - left + 1
	}
	return ans
}
