package main

// https://space.bilibili.com/206214
func countGoodSubarrays(nums []int) (ans int64) {
	type pair struct{ or, left int } // 子数组或值，最小左端点
	orLeft := []pair{}
	last := map[int]int{}

	for i, x := range nums {
		last[x] = i

		// 计算以 i 为右端点的子数组或值
		for j := range orLeft {
			orLeft[j].or |= x
		}
		// x 单独一个数作为子数组
		orLeft = append(orLeft, pair{x, i})

		// 原地去重（相同或值只保留最左边的）
		// 原理见力扣 26. 删除有序数组中的重复项
		idx := 1
		for j := 1; j < len(orLeft); j++ {
			if orLeft[j].or != orLeft[j-1].or {
				orLeft[idx] = orLeft[j]
				idx++
			}
		}
		orLeft = orLeft[:idx]

		for k, p := range orLeft {
			orVal := p.or
			left := p.left
			right := i
			if k < len(orLeft)-1 {
				right = orLeft[k+1].left - 1
			}
			// 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 orVal
			j, ok := last[orVal]
			if ok && j >= left {
				ans += int64(min(right, j) - left + 1)
			}
		}
	}

	return
}
