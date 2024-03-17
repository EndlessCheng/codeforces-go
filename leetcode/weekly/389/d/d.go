package main

import "math"

// https://space.bilibili.com/206214
func minimumMoves(nums []int, k, maxChanges int) int64 {
	pos := []int{}
	c := 0 // nums 中连续的 1 长度
	for i, x := range nums {
		if x == 0 {
			continue
		}
		pos = append(pos, i) // 记录 1 的位置
		c = max(c, 1)
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				c = 3 // 有 3 个连续的 1
			} else {
				c = max(c, 2) // 有 2 个连续的 1
			}
		}
	}

	c = min(c, k)
	if maxChanges >= k-c {
		// 其余 k-c 个 1 可以全部用两次操作得到
		return int64(max(c-1, 0) + (k-c)*2)
	}

	n := len(pos)
	sum := make([]int, n+1)
	for i, x := range pos {
		sum[i+1] = sum[i] + x
	}

	ans := math.MaxInt
	// 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
	size := k - maxChanges
	for right := size; right <= n; right++ {
		// s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 pos[(left+right)/2] 的距离之和
		left := right - size
		i := left + size/2
		s1 := pos[i]*(i-left) - (sum[i] - sum[left])
		s2 := sum[right] - sum[i] - pos[i]*(right-i)
		ans = min(ans, s1+s2)
	}
	return int64(ans + maxChanges*2)
}
