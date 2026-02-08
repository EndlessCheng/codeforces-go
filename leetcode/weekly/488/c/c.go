package main

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k int64) (ans int64) {
	var minQ, maxQ []int
	left := 0
	for right, x := range nums {
		// 1. 入：元素进入窗口
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)

		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// 2. 出：如果窗口不满足要求，左端点离开窗口
		// 只需改下面这行代码，其余逻辑和 2762 题完全一致
		for int64(nums[maxQ[0]]-nums[minQ[0]])*int64(right-left+1) > k {
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// 3. 更新答案
		ans += int64(right - left + 1)
	}
	return
}
