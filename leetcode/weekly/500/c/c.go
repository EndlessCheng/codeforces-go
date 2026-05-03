package main

// https://space.bilibili.com/206214
func minCost(nums []int, queries [][]int) []int {
	n := len(nums)
	sumL := make([]int, n) // sumL[i] 等于从 i 移动到 0 的代价和
	sumR := make([]int, n) // sumR[i] 等于从 0 移动到 i 的代价和
	for i := 1; i < n; i++ {
		// 往左走 i -> i-1
		cost := 1
		if i < n-1 && nums[i]-nums[i-1] > nums[i+1]-nums[i] { // closest(i) = i+1
			cost = nums[i] - nums[i-1] // 只能用方式一往左走
		}
		sumL[i] = sumL[i-1] + cost

		// 往右走 i-1 -> i
		cost = 1
		if i > 1 && nums[i-1]-nums[i-2] <= nums[i]-nums[i-1] { // closest(i-1) = i-2
			cost = nums[i] - nums[i-1] // 只能用方式一往右走
		}
		sumR[i] = sumR[i-1] + cost
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		if l < r {
			// cost(0 -> r) - cost(0 -> l) = cost(l -> r)
			ans[i] = sumR[r] - sumR[l]
		} else {
			// cost(l -> 0) - cost(r -> 0) = cost(l -> r)
			ans[i] = sumL[l] - sumL[r]
		}
	}
	return ans
}
