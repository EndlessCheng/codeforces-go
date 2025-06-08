package main

// https://space.bilibili.com/206214
func countPartitions(nums []int, k int) int {
	const mod = 1_000_000_007
	n := len(nums)
	var minQ, maxQ []int
	f := make([]int, n+1)
	f[0] = 1
	sumF := 0 // 窗口中的 f[i] 之和
	left := 0

	for i, x := range nums {
		// 1. 入
		sumF += f[i]

		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, i)

		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)

		// 2. 出
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			sumF -= f[left]
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// 3. 更新答案
		f[i+1] = sumF % mod
	}
	return f[n]
}
