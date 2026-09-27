package main

// https://space.bilibili.com/206214
// 类似 974. 和可被 K 整除的子数组
func longestSubarrayDivByK(nums []int, k int) (res int) {
	firstPos := map[int]int{0: -1} // 前缀和 % k 首次出现的下标
	sum := 0 // 前缀和
	for r, x := range nums {
		sum = (sum + x%k + k) % k // 保证 sum 非负
		l, ok := firstPos[sum]
		if ok {
			res = max(res, r-l)
		} else {
			firstPos[sum] = r
		}
	}
	return
}

func longestSubarray(nums []int, k int) int {
	// 不取反
	ans := longestSubarrayDivByK(nums, k)

	// 枚举取反元素
	for i := range nums {
		nums[i] *= -1 // 取反
		ans = max(ans, longestSubarrayDivByK(nums, k))
		nums[i] *= -1 // 复原
	}

	return ans
}
