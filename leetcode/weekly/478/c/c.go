package main

// https://space.bilibili.com/206214
func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	ans := n
	lastIndex := make(map[int]int, n) // 预分配空间

	for j, x := range nums {
		if i, ok := lastIndex[x]; ok {
			ans = min(ans, j-i)
		}

		// 计算 reverse(x)
		rev := 0
		for ; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		lastIndex[rev] = j
	}

	if ans == n {
		return -1
	}
	return ans
}
